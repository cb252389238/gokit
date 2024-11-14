package oriRedis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"ori/core/oriConfig"
	"strconv"
	"sync"
	"time"
)

var (
	once sync.Once
	sets *RedisSets
)

type RedisSets struct {
	redis map[string]*redis.Client
	l     sync.RWMutex
}

func (r *RedisSets) Rds(key ...string) *redis.Client {
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

// redis并发锁
func (r *RedisSets) Lock(ctx context.Context, redisName string, key string, lockTime time.Duration) (success bool, unlock func(), err error) {
	value := time.Now().UnixMicro()
	val := strconv.FormatInt(value, 10)
	redisCli := r.Rds(redisName)
	stm, err := redisCli.SetNX(ctx, key, val, lockTime).Result()
	if err != nil {
		return false, nil, err
	}
	if stm == false {
		return false, nil, nil
	}
	return true, func() {
		c := context.Background()
		result, err := redisCli.Get(c, key).Result()
		if err != nil {
			return
		}
		if result != val {
			return
		}
		redisCli.Del(c, key)
	}, nil
}

func New() *RedisSets {
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
			_, err := client.Ping(context.Background()).Result()
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
