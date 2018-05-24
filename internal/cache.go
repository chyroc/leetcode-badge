package internal

import (
	"sync"

	"github.com/patrickmn/go-cache"
)

var defaultCache *cache.Cache
var once = new(sync.Once)

func cacheGet(key string) (*LeetcodeData, bool) {
	once.Do(func() {
		if Conf.Cache {
			defaultCache = cache.New(Conf.CacheTTL, Conf.CacheTTL*2)
		}
	})

	if defaultCache == nil {
		return nil, false
	}
	data, ok := defaultCache.Get(key)
	if !ok {
		return nil, false
	}

	r, ok := data.(*LeetcodeData)
	if !ok || r == nil {
		return nil, false
	}

	return r, true
}

func cacheSet(key string, r *LeetcodeData) {
	once.Do(func() {
		if Conf.Cache {
			defaultCache = cache.New(Conf.CacheTTL, Conf.CacheTTL*2)
		}
	})

	if defaultCache != nil {
		defaultCache.Set(key, r, Conf.CacheTTL)
	}
}
