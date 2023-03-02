package cache

import (
	"sync"
	"time"
)

var (
	once       sync.Once
	speedCache *Cache
)

func New() *Cache {
	once.Do(func() {
		speedCache = NewCache(NoExpiration, time.Second*10)
	})
	return speedCache
}
