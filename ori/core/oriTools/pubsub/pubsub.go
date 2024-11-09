// 发布订阅模型实现
package pubsub

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"
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
	snowflake     *Node
	snLen         int64 //详细补偿队列
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
func New(platform []string) (*Publisher, error) {
	if len(platform) == 0 {
		return nil, errors.New("platform is empty")
	}
	publisherOnce.Do(func() {
		node, _ := NewNode(int64(1))
		publisher = &Publisher{
			m:             sync.RWMutex{},
			buffer:        100,
			subscribers:   make(map[topicType]map[uidType]*Subscriber),
			users:         make(map[uidType]*Subscriber),
			snowflake:     node,
			snLen:         1000,
			platformArray: platform,
		}
	})
	return publisher, nil
}

func getUserId(uid, platform string) string {
	return fmt.Sprintf("%s:%s", uid, platform)
}

func inArray(needle any, haystack any) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	}
	return false
}

// 订阅主题 topic主题 userId 用户id platform客户端类型
func (p *Publisher) Subscribe(topic, uid, platform string) (*Subscriber, error) {
	p.m.Lock()
	defer p.m.Unlock()
	if uid == "" {
		return nil, errors.New("uid is empty")
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
	userId := getUserId(uid, platform)
	//用户是否已经创建订阅实体
	if val, ok := p.users[userId]; ok {
		sub = val
		sub.Topic[topic] = struct{}{}
	} else {
		sub = &Subscriber{
			C:        make(chan any, p.buffer), //消息接受通道
			Topic:    map[topicType]struct{}{topic: {}},
			uid:      userId,
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
	p.users[userId] = sub
	return sub, nil
}

// 退出订阅
func (p *Publisher) Exit(topic, uid, platform string) error {
	if uid == "" {
		return errors.New("uid is empty")
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
	userId := getUserId(uid, platform)
	p.m.Lock()
	defer p.m.Unlock()
	if subSets, ok := p.subscribers[topic]; ok {
		delete(subSets, userId)
	}
	if user, ok := p.users[userId]; ok {
		delete(user.Topic, topic)
	}
	return nil
}

// 发布一个主题信息
func (p *Publisher) Publish(topic string, message any, platform ...string) {
	p.m.RLock()
	defer p.m.RUnlock()
	if subSets, ok := p.subscribers[topic]; ok {
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
func (p *Publisher) PublishToUser(uid string, message any, platform ...string) {
	p.m.RLock()
	defer p.m.RUnlock()
	if len(platform) > 0 {
		for _, plat := range platform {
			if !inArray(plat, p.platformArray) {
				continue
			}
			userId := getUserId(uid, plat)
			if sub, ok := p.users[userId]; ok {
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
			userId := getUserId(uid, plat)
			if sub, ok := p.users[userId]; ok {
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
func (p *Publisher) Replier(uid, platform string, sn int64) {
	p.m.RLock()
	defer p.m.RUnlock()
	userId := getUserId(uid, platform)
	if sub, ok := p.users[userId]; ok {
		offset := sn % p.snLen
		buildMsg := sub.store[offset]
		if time.Now().Unix()-buildMsg.createTime > 60 {
			return
		}
		buildMsgByte, _ := json.Marshal(buildMsg)
		sub.C <- string(buildMsgByte)
	}
}
