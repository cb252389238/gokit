package svc

import (
	"context"
	"frame/internal/factory"
	"frame/resource"
	"frame/util/pool"
	"os"
	"sync"
)

type ServiceContext struct {
	Wg      *sync.WaitGroup
	C       chan os.Signal //全局信号
	WsC     chan os.Signal //websocket信号
	HttpC   chan os.Signal
	L       *sync.RWMutex
	Context context.Context
	Cancle  context.CancelFunc
	Mysql   *resource.MysqlSets
	Redis   *resource.RedisSets
	Pool    pool.Pool        //通用连接池
	Factory *factory.Factory //工厂类
}

func NewServiceContext() *ServiceContext {
	cancel, cancelFunc := context.WithCancel(context.Background())
	gmPool := resource.NewPool(
		func() (interface{}, error) {
			return 1, nil
		},
		func(v interface{}) error {
			return nil
		},
		100,
		100,
		1000,
	)
	ctx := &ServiceContext{
		Wg:      &sync.WaitGroup{},
		C:       make(chan os.Signal),
		WsC:     make(chan os.Signal),
		L:       &sync.RWMutex{},
		Context: cancel,
		Cancle:  cancelFunc,
		//Mysql:     resource.NewDb(),
		//Redis: resource.NewRedis(),
		Pool:    gmPool,
		Factory: factory.New(),
	}
	return ctx
}
