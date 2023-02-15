package oriEngine

import (
	"context"
	"ori/internal/core/database"
	"ori/internal/core/log"
	"ori/internal/core/pool"
	"ori/internal/core/redis"
	"ori/internal/factory"
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
	Mysql      *database.MysqlSets
	Redis      *redis.RedisSets
	Pool       pool.Pool        //通用连接池
	Factory    *factory.Factory //工厂类
	Log        *log.LocalLogger
}

func NewOriEngine() *OriEngine {
	cancel, cancelFunc := context.WithCancel(context.Background())
	ctx := &OriEngine{
		Wg:         &sync.WaitGroup{},
		Signal:     make(chan os.Signal),
		WsSignal:   make(chan os.Signal),
		HttpSignal: make(chan os.Signal),
		L:          &sync.RWMutex{},
		Context:    cancel,
		Cancel:     cancelFunc,
		//Mysql:      database.NewDb(),
		//Redis:      redis.NewRedis(),
		Pool: pool.NewPool(
			func() (interface{}, error) {
				return 1, nil
			},
			func(v interface{}) error {
				return nil
			},
			100,
			100,
			1000,
		),
		//Factory: factory.New(),
		Log: log.NewLog(),
	}
	return ctx
}
