package oriPool

import (
	"sync"
	"time"
)

var (
	pool     Pool
	poolOnce sync.Once
	err      error
)

func New(open func() (interface{}, error), close func(interface{}) error, cap, maxIdle, maxCap int) (Pool, error) {
	poolOnce.Do(func() {
		poolConfig := &Config{
			InitialCap: cap,     //资源池初始连接数
			MaxIdle:    maxIdle, //最大空闲连接数
			MaxCap:     maxCap,  //最大并发连接数
			Factory:    open,
			Close:      close,
			//Ping:       ping,
			//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
			IdleTimeout: 15 * time.Second,
		}
		pool, err = NewChannelPool(poolConfig)
	})
	return pool, err
}
