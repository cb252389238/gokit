// 发布订阅模型实现
package centrifuge

import (
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"time"
)

const (
	TopicUserOnline = "user_online"
)

var publisher *Publisher
var publisherOnce sync.Once

type topicType = string
type uidType = string
type platformType = string //客户端类型

// 发布者
type Publisher struct {
	m             sync.RWMutex                          //读写锁
	buffer        int                                   //订阅队列缓存大小
	subscribers   map[topicType]map[uidType]*Subscriber //主题订阅者集合
	users         map[uidType]*Subscriber               //用户和订阅实例绑定关系
	snLen         int64                                 //详细补偿队列
	platformArray []string
}

type snMsg struct {
	Sn         int64 `json:"sn"` //消息序列号
	Msg        any   `json:"msg"`
	createTime int64 `json:"-"`
}

// 订阅者
type Subscriber struct {
	C        chan any
	Topic    map[topicType]struct{} //订阅得主题
	uid      uidType                //用户标识
	store    []snMsg                //消息缓存
	sn       int64                  //消息序列号
	platform platformType           //客户端类型
}

// 构建一个新的发布者对象
func NewPublisher(platform []string) (*Publisher, error) {
	if len(platform) == 0 {
		return nil, errors.New("platform is empty")
	}
	publisherOnce.Do(func() {
		publisher = &Publisher{
			m:             sync.RWMutex{},
			buffer:        100,
			subscribers:   make(map[topicType]map[uidType]*Subscriber),
			users:         make(map[uidType]*Subscriber),
			snLen:         1000,
			platformArray: platform,
		}
	})
	return publisher, nil
}

func inArray(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// 订阅主题 topic主题 userId 用户id platform客户端类型
func (p *Publisher) Subscribe(topic, userId, platform string) (*Subscriber, error) {
	p.m.Lock()
	defer p.m.Unlock()
	if userId == "" {
		return nil, errors.New("userId is empty")
	}
	if platform == "" {
		return nil, errors.New("platform is empty")
	}
	if topic == "" {
		return nil, errors.New("topic is empty")
	}
	if !inArray(platform, p.platformArray) {
		return nil, errors.New("platform is not exist")
	}
	var sub *Subscriber
	uid := getUid(userId, platform)
	//用户是否已经创建订阅实体
	if val, ok := p.users[uid]; ok {
		sub = val
		sub.Topic[topic] = struct{}{}
	} else {
		sub = &Subscriber{
			C:        make(chan any, p.buffer), //消息接受通道
			Topic:    map[topicType]struct{}{topic: {}},
			uid:      uid,
			store:    make([]snMsg, p.snLen),
			platform: platform,
		}
	}
	if subSets, ok := p.subscribers[topic]; ok {
		if s, ok := subSets[sub.uid]; ok {
			s.Topic[topic] = struct{}{}
		} else {
			subSets[sub.uid] = sub
		}
	} else {
		subSets = map[uidType]*Subscriber{
			sub.uid: sub,
		}
		p.subscribers[topic] = subSets
	}
	p.users[uid] = sub
	return sub, nil
}

// 退出订阅
func (p *Publisher) UnSubscribe(topic, userId, platform string) error {
	if userId == "" {
		return errors.New("userId is empty")
	}
	if platform == "" {
		return errors.New("platform is empty")
	}
	if topic == "" {
		return errors.New("topic is empty")
	}
	if !inArray(platform, p.platformArray) {
		return errors.New("platform is not exist")
	}
	uid := getUid(userId, platform)
	p.m.Lock()
	defer p.m.Unlock()
	if subSets, ok := p.subscribers[topic]; ok {
		delete(subSets, uid)
	}
	if user, ok := p.users[uid]; ok {
		delete(user.Topic, topic)
	}
	return nil
}

// 发布一个主题信息
func (p *Publisher) Publish(topic string, message any, platform ...string) {
	p.m.RLock()
	defer p.m.RUnlock()
	subSets, ok := p.subscribers[topic]
	if ok {
		for _, v := range subSets {
			split := strings.Split(v.uid, ":")
			if len(split) != 2 {
				continue
			}
			if len(platform) > 0 && !inArray(split[1], platform) {
				continue
			}
			offset := v.sn % p.snLen
			buildMsg := snMsg{
				Sn:         v.sn,
				Msg:        message,
				createTime: time.Now().Unix(),
			}
			v.store[offset] = buildMsg
			buildMsgByte, _ := json.Marshal(buildMsg)
			sub := v
			select {
			case sub.C <- string(buildMsgByte):
			default:
				//忽略数据
			}
			v.sn++
		}
	}
}

// 向所有用户推送主题
func (p *Publisher) PublishAll(message any, platform ...string) {
	p.m.RLock()
	defer p.m.RUnlock()
	for _, v := range p.users {
		split := strings.Split(v.uid, ":")
		if len(split) != 2 {
			continue
		}
		if len(platform) > 0 && !inArray(split[1], platform) {
			continue
		}
		offset := v.sn % p.snLen
		buildMsg := snMsg{
			Sn:         v.sn,
			Msg:        message,
			createTime: time.Now().Unix(),
		}
		v.store[offset] = buildMsg
		buildMsgByte, _ := json.Marshal(buildMsg)
		sub := v
		select {
		case sub.C <- string(buildMsgByte):
		default:
			//忽略数据
		}
		v.sn++
	}
}

// 向单个用户发送主题
func (p *Publisher) PublishToUser(userId string, message any, platform ...string) {
	p.m.RLock()
	defer p.m.RUnlock()
	if len(platform) > 0 {
		for _, plat := range platform {
			if !inArray(plat, p.platformArray) {
				continue
			}
			uid := getUid(userId, plat)
			if sub, ok := p.users[uid]; ok {
				offset := sub.sn % p.snLen
				buildMsg := snMsg{
					Sn:         sub.sn,
					Msg:        message,
					createTime: time.Now().Unix(),
				}
				sub.store[offset] = buildMsg
				buildMsgByte, _ := json.Marshal(buildMsg)
				sub.C <- string(buildMsgByte)
				sub.sn++
			}
		}
	} else {
		for _, plat := range p.platformArray {
			uid := getUid(userId, plat)
			if sub, ok := p.users[uid]; ok {
				offset := sub.sn % p.snLen
				buildMsg := snMsg{
					Sn:         sub.sn,
					Msg:        message,
					createTime: time.Now().Unix(),
				}
				sub.store[offset] = buildMsg
				buildMsgByte, _ := json.Marshal(buildMsg)
				sub.C <- string(buildMsgByte)
				sub.sn++
			}
		}
	}
}

// 发送补偿消息
func (p *Publisher) Replier(userId, platform string, sn []int64) {
	p.m.RLock()
	defer p.m.RUnlock()
	uid := getUid(userId, platform)
	if sub, ok := p.users[uid]; ok {
		for _, v := range sn {
			offset := v % p.snLen
			buildMsg := sub.store[offset]
			if time.Now().Unix()-buildMsg.createTime > 60 {
				return
			}
			buildMsgByte, _ := json.Marshal(buildMsg)
			sub.C <- string(buildMsgByte)
		}
	}
}
