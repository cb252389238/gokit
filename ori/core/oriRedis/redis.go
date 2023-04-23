package oriRedis

import (
	"fmt"
	"github.com/spf13/cast"
	"ori/core/oriConfig"
	"sync"

	"github.com/go-redis/redis"
)

var (
	once sync.Once
	sets *RedisSets
)

type RedisSets struct {
	redis map[string]*redis.Client
	l     sync.RWMutex
}

func (r *RedisSets) Redis(key ...string) *redis.Client {
	r.l.RLock()
	defer r.l.RUnlock()
	name := "default"
	if len(key) > 0 {
		name = key[0]
	}
	if client, ok := r.redis[name]; ok {
		return client
	}
	return nil
}

func NewRedis() *RedisSets {
	if !oriConfig.GetHotConf().Services.REDIS_SERVER {
		return nil
	}
	once.Do(func() {
		conf := oriConfig.GetHotConf()
		redisSets := map[string]*redis.Client{}
		for _, r := range conf.Redis {
			client := redis.NewClient(&redis.Options{
				Addr:       fmt.Sprintf("%s:%s", r.Host, r.Port),
				Password:   r.Password,
				DB:         cast.ToInt(r.Database),
				MaxRetries: 3, //重试次数
			})
			_, err := client.Ping().Result()
			if err != nil {
				panic(err)
			}
			redisSets[r.Name] = client
		}
		sets = &RedisSets{
			redis: redisSets,
		}
	})
	return sets
}

func Redis(keys ...string) *redis.Client {
	rds := NewRedis()
	key := "default"
	if len(keys) > 0 {
		key = keys[0]
	}
	return rds.Redis(key)
}
