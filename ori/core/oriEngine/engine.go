package oriEngine

import (
	"context"
	"github.com/blinkbean/dingtalk"
	cache2 "ori/core/oriCache"
	"ori/core/oriConfig"
	"ori/core/oriDb"
	log2 "ori/core/oriLog"
	pool2 "ori/core/oriPool"
	"ori/core/oriRedis"
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
	Mysql      *oriDb.MysqlSets
	Redis      *oriRedis.RedisSets
	Pool       pool2.Pool //通用连接池
	Log        *log2.LocalLogger
	Cache      *cache2.Cache
	WebHook    *dingtalk.DingTalk
}

func NewOriEngine() *OriEngine {
	webHookCli := dingtalk.InitDingTalkWithSecret(oriConfig.GetHotConf().WebHookToken, oriConfig.GetHotConf().WebHookSecret)
	cancel, cancelFunc := context.WithCancel(context.Background())
	ctx := &OriEngine{
		Wg:         &sync.WaitGroup{},
		Signal:     make(chan os.Signal),
		WsSignal:   make(chan os.Signal),
		HttpSignal: make(chan os.Signal),
		L:          &sync.RWMutex{},
		Context:    cancel,
		Cancel:     cancelFunc,
		Mysql:      oriDb.NewDb(),
		Redis:      oriRedis.NewRedis(),
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
