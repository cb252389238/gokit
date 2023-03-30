package rateLimit

import (
	"ori/core/oriCache"
	"time"
)

type RateLimit struct {
	Cache     *oriCache.Cache
	LifeSpan  int
	Frequency int
}

func New(lifeSpan, frequency int) *RateLimit {
	return &RateLimit{
		Cache:     oriCache.NewCache(time.Second*time.Duration(lifeSpan), time.Minute*10),
		LifeSpan:  lifeSpan,
		Frequency: frequency,
	}
}

// true可访问 false不可访问
func (r *RateLimit) CheckRate(key string) bool {
	i, ok := r.Cache.Get(key)
	if !ok {
		r.Cache.Set(key, 1, oriCache.DefaultExpiration)
		return true
	}
	num := i.(int)
	if num < r.Frequency {
		r.Cache.IncrementInt(key, 1)
		return true
	}
	return false
}
