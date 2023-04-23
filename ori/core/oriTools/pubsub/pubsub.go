// 发布订阅模型实现
package pubsub

import (
	"sync"
)

// 发布者对象
type Publisher struct {
	m           sync.RWMutex                     //读写锁
	buffer      int                              //订阅队列缓存大小
	subscribers map[string]map[int64]*Subscriber //频道订阅者信息
	users       map[string]*Subscriber           //单个用户信息
	snowflake   *Node
}

type Subscriber struct {
	C     chan interface{}
	Key   int64
	topic string
	uid   string
}

// 构建一个新的发布者对象
func NewPublisher(buffer int) (*Publisher, error) {
	node, err := NewNode(int64(1))
	if err != nil {
		return nil, err
	}
	return &Publisher{
		m:           sync.RWMutex{},
		buffer:      buffer,
		subscribers: make(map[string]map[int64]*Subscriber),
		snowflake:   node,
	}, nil
}

// 添加一个新的订阅者
func (p *Publisher) SubscriberTopic(topic, uid string) *Subscriber {
	p.m.Lock()
	defer p.m.Unlock()
	var sub *Subscriber
	key := p.snowflake.Generate().Int64()
	sub = &Subscriber{
		Key:   key,
		C:     make(chan interface{}, p.buffer),
		topic: topic,
		uid:   uid,
	}
	if subSets, ok := p.subscribers[topic]; ok {
		subSets[key] = sub
	} else {
		subSets = map[int64]*Subscriber{
			key: sub,
		}
		p.subscribers[topic] = subSets
	}
	p.users[uid] = sub
	return sub
}

// 退出订阅
func (p *Publisher) Exit(sub *Subscriber) {
	if sub == nil {
		return
	}
	p.m.Lock()
	defer p.m.Unlock()
	if subSets, ok := p.subscribers[sub.topic]; ok {
		delete(subSets, sub.Key)
		delete(p.users, sub.uid)
	}
}

/*func (p *Publisher) Del(sub *Subscriber) {
	if sub == nil {
		return
	}
	p.m.Lock()
	defer p.m.Unlock()
	if subSets, ok := p.subscribers[sub.topic]; ok {
		delete(subSets, sub.Key)
	}
}*/

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
func (p *Publisher) PublishAll(message interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()
	for _, subscriberTopic := range p.subscribers {
		for _, v := range subscriberTopic {
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

// 向单个用户发送主题
func (p *Publisher) PublishToUser(uid string, message interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()
	if sub, ok := p.users[uid]; ok {
		sub.C <- message
	}
}
