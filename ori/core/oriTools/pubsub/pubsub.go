// 发布订阅模型实现
package pubsub

import (
	"errors"
	"sync"
)

const (
	DEFAULT_TOPIC = "default_topic"
)

type topicType = string
type uidType = string
type keyType = int64

// 发布者
type Publisher struct {
	m           sync.RWMutex                          //读写锁
	buffer      int                                   //订阅队列缓存大小
	subscribers map[topicType]map[keyType]*Subscriber //主题订阅者集合
	users       map[uidType]*Subscriber               //用户和订阅实例绑定关系
	snowflake   *Node
}

type Subscriber struct {
	C     chan any
	Key   keyType
	Topic map[topicType]struct{} //订阅得主题
	uid   uidType
}

// 构建一个新的发布者对象
func New() *Publisher {
	node, _ := NewNode(int64(1))
	return &Publisher{
		m:           sync.RWMutex{},
		buffer:      10,
		subscribers: make(map[topicType]map[keyType]*Subscriber),
		users:       make(map[uidType]*Subscriber),
		snowflake:   node,
	}
}

// 订阅一个主题
func (p *Publisher) Subscribe(topic, uid string) (*Subscriber, error) {
	if uid == "" {
		return nil, errors.New("uid is empty")
	}
	if topic == "" {
		topic = DEFAULT_TOPIC
	}
	p.m.Lock()
	defer p.m.Unlock()
	var sub *Subscriber

	//用户是否已经创建订阅实体
	if val, ok := p.users[uid]; ok {
		sub = val
		sub.Topic[topic] = struct{}{}
	} else {
		key := p.snowflake.Generate().Int64()
		sub = &Subscriber{
			C:     make(chan any, p.buffer), //消息接受通道
			Topic: map[topicType]struct{}{topic: {}},
			uid:   uid,
			Key:   key,
		}
	}

	if subSets, ok := p.subscribers[topic]; ok {
		if s, ok := subSets[sub.Key]; ok {
			s.Topic[topic] = struct{}{}
		} else {
			subSets[sub.Key] = sub
		}
	} else {
		subSets = map[int64]*Subscriber{
			sub.Key: sub,
		}
		p.subscribers[topic] = subSets
	}
	p.users[uid] = sub
	return sub, nil
}

// 退出订阅
func (p *Publisher) Exit(sub *Subscriber, topic string) {
	if sub == nil {
		return
	}
	p.m.Lock()
	defer p.m.Unlock()
	if subSets, ok := p.subscribers[topic]; ok {
		delete(subSets, sub.Key)
	}
	if user, ok := p.users[sub.uid]; ok {
		delete(user.Topic, topic)
	}
}

// 发布一个主题信息
func (p *Publisher) Publish(topic string, message string) {
	p.m.RLock()
	defer p.m.RUnlock()
	if subSets, ok := p.subscribers[topic]; ok {
		for _, v := range subSets {
			go func(v *Subscriber) {
				select {
				case v.C <- message:
				default:
					//忽略数据
				}
			}(v)
		}
	}
}

// 向所有房间推送主题
func (p *Publisher) PublishAll(message any) {
	p.m.RLock()
	defer p.m.RUnlock()
	for _, sub := range p.users {
		go func(v *Subscriber) {
			select {
			case v.C <- message:
			default:
				//忽略数据
			}
		}(sub)
	}
}

// 向单个用户发送主题
func (p *Publisher) PublishToUser(uid string, message any) {
	p.m.RLock()
	defer p.m.RUnlock()
	if sub, ok := p.users[uid]; ok {
		sub.C <- message
	}
}
