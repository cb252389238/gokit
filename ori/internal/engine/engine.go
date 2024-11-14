package engine

import (
	"context"
	cache2 "ori/core/oriCache"
	"ori/core/oriConfig"
	"ori/core/oriDb"
	"ori/core/oriLog"
	"ori/core/oriPool"
	"ori/core/oriRedis"
	"ori/core/oriSnowflake"
	"ori/core/oriTools/cache"
	"ori/core/oriTools/easy"
	"os"
	"sync"
)

type OriEngine struct {
	Wg         *sync.WaitGroup
	Signal     chan os.Signal //全局控制信号
	WsSignal   chan os.Signal //websocket信号
	HttpSignal chan os.Signal
	L          *sync.RWMutex
	Context    context.Context
	Cancel     context.CancelFunc
	Cache      *cache.Cache
}

func NewOriEngine() *OriEngine {
	oriLog.NewLog() //初始化日志
	//初始化上下文
	cancel, cancelFunc := context.WithCancel(context.Background())
	//初始化redis
	if len(oriConfig.GetHotConf().Redis) >= 1 {
		oriRedis.New()
	}
	if len(oriConfig.GetHotConf().Mysql) >= 1 {
		oriDb.New()
	}
	ip := easy.GetLocalIp()
	intIp := easy.Ipv4StringToInt(ip)
	node := intIp % 1000
	err := oriSnowflake.New(node)
	if err != nil {
		panic(err)
	}
	_, err = oriPool.New(
		func() (interface{}, error) {
			return 1, nil
		},
		func(v interface{}) error {
			return nil
		},
		10,  //初始化连接数
		50,  //最大空闲连接数
		500, //最大并发连接数
	)
	if err != nil {
		panic(err)
	}
	ctx := &OriEngine{
		Wg:         &sync.WaitGroup{},
		Signal:     make(chan os.Signal),
		WsSignal:   make(chan os.Signal),
		HttpSignal: make(chan os.Signal),
		L:          &sync.RWMutex{},
		Context:    cancel,
		Cancel:     cancelFunc,
		Cache:      cache2.New(),
	}
	return ctx
}
