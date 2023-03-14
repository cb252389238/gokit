package rateLimit

import (
	"ori/internal/core/cache"
	"time"
)

type RateLimit struct {
	Cache     *cache.Cache
	LifeSpan  int
	Frequency int
}

func New(lifeSpan, frequency int) *RateLimit {
	return &RateLimit{
		Cache:     cache.NewCache(time.Second*time.Duration(lifeSpan), time.Minute*10),
		LifeSpan:  lifeSpan,
		Frequency: frequency,
	}
}

// true可访问 false不可访问
func (r *RateLimit) CheckRate(key string) bool {
	i, ok := r.Cache.Get(key)
	if !ok {
		r.Cache.Set(key, 1, cache.DefaultExpiration)
		return true
	}
	num := i.(int)
	if num < r.Frequency {
		r.Cache.IncrementInt(key, 1)
		return true
	}
	return false
}
