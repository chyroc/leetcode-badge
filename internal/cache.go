package internal

import (
	"github.com/patrickmn/go-cache"
)

var defaultCache *cache.Cache

func init() {
	if Conf.Cache {
		defaultCache = cache.New(Conf.CacheTTL, Conf.CacheTTL*2)
	}
}

func cacheGet(key string) (*LeetcodeData, bool) {
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
	if defaultCache != nil {
		defaultCache.Set(key, r, Conf.CacheTTL)
	}
}
