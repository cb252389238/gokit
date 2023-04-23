package oriCache

import (
	"ori/core/oriTools/cache"
	"sync"
	"time"
)

var (
	once       sync.Once
	speedCache *cache.Cache
)

func New() *cache.Cache {
	once.Do(func() {
		speedCache = cache.New(cache.NoExpiration, time.Second*10)
	})
	return speedCache
}
