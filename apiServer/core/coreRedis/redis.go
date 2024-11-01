package coreRedis

import (
	"apiServer/core/coreConfig"
	error2 "apiServer/i18n/error"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"strconv"
	"sync"
	"time"
)

var (
	once sync.Once
	sets *RedisSets
)

const (
	USER_INSTANCE_NAME     = "user"     //用户redis
	IM_INSTANCE_NAME       = "im"       //list redis
	CHATROOM_INSTANCE_NAME = "chatroom" //房间redis
)

type RedisSets struct {
	redis map[string]*redis.Client
	l     sync.RWMutex
}

func (r *RedisSets) getRedis(key ...string) *redis.Client {
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
	once.Do(func() {
		conf := coreConfig.GetHotConf()
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
				panic("redis初始化失败:" + err.Error())
			}
			redisSets[r.Name] = client
		}
		sets = &RedisSets{
			redis: redisSets,
		}
	})
	return sets
}

func getRedis(keys ...string) *redis.Client {
	rds := NewRedis()
	key := ""
	if len(keys) > 0 {
		key = keys[0]
	}
	return rds.getRedis(key)
}

func (r *RedisSets) GetUserRedis() *redis.Client {
	return getRedis(USER_INSTANCE_NAME)
}

func (r *RedisSets) UserLock(ctx context.Context, key string, lockTime time.Duration) (success bool, unlock func(), err error) {
	redisCli := r.getRedis(USER_INSTANCE_NAME)
	return baseLock(ctx, key, lockTime, redisCli)
}

func UserLock(ctx context.Context, key string, lockTime time.Duration) (success bool, unlock func(), err error) {
	redisCli := getRedis(USER_INSTANCE_NAME)
	return baseLock(ctx, key, lockTime, redisCli)
}

func baseLock(ctx context.Context, key string, lockTime time.Duration, redisCli *redis.Client) (success bool, unlock func(), err error) {
	maxTryNum := 2000
	tryTime := 3 * time.Millisecond
	value := time.Now().UnixMicro()
	val := strconv.FormatInt(value, 10)
	stm := false
	for i := 0; i < maxTryNum; i++ {
		if stm {
			break
		}
		stm, err = redisCli.SetNX(ctx, key, val, lockTime).Result()
		if err != nil {
			panic(error2.I18nError{
				Code: error2.ErrorCodeSystemBusy,
				Msg:  nil,
			})
		}
		time.Sleep(tryTime)
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
