package resource

import (
	"fmt"
	"github.com/spf13/cast"
	"sync"

	"github.com/go-redis/redis"
)

var (
	redisOnce sync.Once
	rsets     *RedisSets
)

type RedisSets struct {
	redis map[string]*redis.Client
	l     sync.RWMutex
}

func (r *RedisSets) Key(key ...string) *redis.Client {
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
	redisOnce.Do(func() {
		conf := GetAllHotConf()
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
		rsets = &RedisSets{
			redis: redisSets,
		}
	})
	return rsets
}

// GetRedis
//  @Description: 根据redis配置别名获取对应的redis单例对象
//  @param keys redis配置别名 不传默认是"shentu"
//  @return *redis.Client
//
func GetRedis(keys ...string) *redis.Client {
	rds := NewRedis()
	key := "default"
	if len(keys) > 0 {
		key = keys[0]
	}
	return rds.Key(key)
}
