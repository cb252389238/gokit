package engine

import (
	"context"
	"github.com/blinkbean/dingtalk"
	cache2 "ori/core/oriCache"
	"ori/core/oriConfig"
	"ori/core/oriDb"
	log2 "ori/core/oriLog"
	pool2 "ori/core/oriPool"
	"ori/core/oriRedis"
	"ori/core/oriTools/cache"
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
	Db         *oriDb.MysqlSets
	Redis      *oriRedis.RedisSets
	Pool       pool2.Pool //通用连接池
	Log        *log2.LocalLogger
	Cache      *cache.Cache
	WebHook    *dingtalk.DingTalk
}

func NewOriEngine() *OriEngine {
	webHookCli := dingtalk.InitDingTalkWithSecret(oriConfig.GetHotConf().WebHookToken, oriConfig.GetHotConf().WebHookSecret)
	cancel, cancelFunc := context.WithCancel(context.Background())
	var redis *oriRedis.RedisSets
	if len(oriConfig.GetHotConf().Redis) >= 1 {
		redis = oriRedis.NewRedis()
	} else {
		redis = nil
	}
	var db *oriDb.MysqlSets
	if len(oriConfig.GetHotConf().Mysql) >= 1 {
		db = oriDb.NewDb()
	} else {
		db = nil
	}
	ctx := &OriEngine{
		Wg:         &sync.WaitGroup{},
		Signal:     make(chan os.Signal),
		WsSignal:   make(chan os.Signal),
		HttpSignal: make(chan os.Signal),
		L:          &sync.RWMutex{},
		Context:    cancel,
		Cancel:     cancelFunc,
		Db:         db,
		Redis:      redis,
		Pool: pool2.NewPool(
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
		Log:     log2.NewLog(),
		Cache:   cache2.New(),
		WebHook: webHookCli,
	}
	return ctx
}
