//发布订阅模型实现
package pubsub

import (
	"fmt"
	"sync"
	"time"
)

//发布者对象
type publisher struct {
	m           sync.RWMutex                     //读写锁
	buffer      int                              //订阅队列缓存大小
	subscribers map[string]map[int64]*Subscriber //订阅者信息
	snowflake   *Node
}

type Subscriber struct {
	C     chan interface{}
	Key   int64
	topic string
}

//构建一个新的发布者对象
func NewPublisher(buffer int) (*publisher, error) {
	node, err := NewNode(int64(1))
	if err != nil {
		return nil, err
	}
	return &publisher{
		m:           sync.RWMutex{},
		buffer:      buffer,
		subscribers: make(map[string]map[int64]*Subscriber),
		snowflake:   node,
	}, nil
}

//添加一个新的订阅者
func (p *publisher) SubscriberTopic(topic string) *Subscriber {
	p.m.Lock()
	defer p.m.Unlock()
	var sub *Subscriber
	key := p.snowflake.Generate().Int64()
	sub = &Subscriber{
		Key:   key,
		C:     make(chan interface{}, p.buffer),
		topic: topic,
	}
	if subSets, ok := p.subscribers[topic]; ok {
		subSets[key] = sub
	} else {
		subSets = map[int64]*Subscriber{
			key: sub,
		}
		p.subscribers[topic] = subSets
	}
	return sub
}

//退出订阅
func (p *publisher) Exit(sub *Subscriber) {
	p.m.Lock()
	defer p.m.Unlock()
	if subSets, ok := p.subscribers[sub.topic]; ok {
		delete(subSets, sub.Key)
		close(sub.C)
	}
}

//发布一个主题信息
func (p *publisher) Publish(topic string, message string) {
	p.m.RLock()
	defer p.m.RUnlock()
	if subSets, ok := p.subscribers[topic]; ok {
		for _, v := range subSets {
			go func(v *Subscriber) {
				select {
				case v.C <- message:
				case <-time.After(time.Millisecond * 5):
					fmt.Println("插入超时")
					<-v.C
					return
				}
			}(v)
		}
	}
}
