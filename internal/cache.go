package internal

import (
	"sync"

	"github.com/patrickmn/go-cache"
)

var defaultCache *cache.Cache
var once = new(sync.Once)

func cacheGetLeetcode(key string) (*LeetcodeData, bool) {
	once.Do(func() {
		if Conf.Cache {
			defaultCache = cache.New(Conf.CacheTTL, Conf.CacheTTL*2)
		}
	})

	if defaultCache == nil {
		return nil, false
	}
	data, ok := defaultCache.Get("Leetcode" + key)
	if !ok {
		return nil, false
	}

	r, ok := data.(*LeetcodeData)
	if !ok || r == nil {
		return nil, false
	}

	return r, true
}

func cacheSetLeetcode(key string, r *LeetcodeData) {
	once.Do(func() {
		if Conf.Cache {
			defaultCache = cache.New(Conf.CacheTTL, Conf.CacheTTL*2)
		}
	})

	if defaultCache != nil {
		defaultCache.Set("Leetcode"+key, r, Conf.CacheTTL)
	}
}

func cacheGetShields(key string) (string, bool) {
	once.Do(func() {
		if Conf.Cache {
			defaultCache = cache.New(Conf.CacheTTL, Conf.CacheTTL*2)
		}
	})

	if defaultCache == nil {
		return "", false
	}
	data, ok := defaultCache.Get("Shields" + key)
	if !ok {
		return "", false
	}

	r, ok := data.(string)
	if !ok || r == "" {
		return "", false
	}

	return r, true
}

func cacheSetShields(key string, r string) {
	once.Do(func() {
		if Conf.Cache {
			defaultCache = cache.New(Conf.CacheTTL, Conf.CacheTTL*2)
		}
	})

	if defaultCache != nil {
		defaultCache.Set("Shields"+key, r, Conf.CacheTTL)
	}
}
