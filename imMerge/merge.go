package imMerge

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Engine struct {
	mergeThreshold int              //合并条目阈值 大于等于阈值则合并
	mergeTime      int64            //合并阈值时间
	timeOut        int64            //协程超时时间 毫秒
	data           chan MessageData //接收数据通道
	Result         chan MessageData //结果通道
	ctx            context.Context
	l              sync.RWMutex
	groupIdPool    *SafeMap
}

type Config struct {
	MergeThreshold int   //合并条目阈值 大于等于阈值则合并
	MergeTime      int64 //合并阈值时间
	TimeOut        int64 //协程超时时间
	Ctx            context.Context
}

type MessageData struct {
	groupId string
	data    any
}

func New(conf Config) *Engine {
	engine := &Engine{
		mergeThreshold: conf.MergeThreshold,
		mergeTime:      conf.MergeTime,
		timeOut:        conf.TimeOut,
		data:           make(chan MessageData, 10000),
		Result:         make(chan MessageData, 10000),
		ctx:            conf.Ctx,
		groupIdPool:    NewSafeMap(),
	}
	go engine.run()
	return engine
}

func (e *Engine) run() {
	defer fmt.Println("imMerge 退出")
	for {
		select {
		case <-e.ctx.Done():
			close(e.Result)
			return
		case data := <-e.data:
			if ch1, ok := e.groupIdPool.Get(data.groupId); !ok {
				e.groupIdPool.Set(data.groupId, make(chan any, 10000))
				if ch2, ok := e.groupIdPool.Get(data.groupId); ok {
					ch2.(chan any) <- data.data
				}
				go e.merge(data.groupId)
			} else {
				ch1.(chan any) <- data.data
			}
		}
	}
}

func (e *Engine) Message(groupId string, data any) {
	m := MessageData{
		groupId: groupId,
		data:    data,
	}
	e.data <- m
}

func (e *Engine) merge(groupId string) {
	defer fmt.Printf("merge: 协程退出,groupId:%s\r\n", groupId)
	msgNum := 0
	mergeData := []any{}
	tickTimeOut := time.NewTicker(time.Millisecond * time.Duration(e.timeOut))
	tickMergeThreshold := time.NewTicker(time.Millisecond * time.Duration(e.mergeTime))
	for {
		chmsg, ok := e.groupIdPool.Get(groupId)
		if !ok {
			return
		}
		select {
		case <-tickTimeOut.C: //规定时间内房间内没有消息发送释放协程 节省内存资源
			if msgNum <= 0 { //如果超时时间内没有消息合并则释放协程
				e.groupIdPool.Del(groupId)
				return
			}
			msgNum = 0
		case <-tickMergeThreshold.C: //规定阈值时间内返回消息
			if len(mergeData) <= 0 {
				continue
			}
			e.Result <- MessageData{
				groupId: groupId,
				data:    mergeData,
			}
			mergeData = []any{}
		case v := <-chmsg.(chan any):
			msgNum++
			if len(mergeData) >= e.mergeThreshold {
				e.Result <- MessageData{
					groupId: groupId,
					data:    mergeData,
				}
				mergeData = []any{}

			}
			mergeData = append(mergeData, v)
		}
	}
}
