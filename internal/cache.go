package internal

import (
	"github.com/patrickmn/go-cache"
)

var defaultCache *cache.Cache

func cacheGetLeetcode(key string) (*LeetcodeData, bool) {
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
	if defaultCache != nil {
		defaultCache.Set("Leetcode"+key, r, Conf.CacheTTL)
	}
}

func cacheGetShields(key string) (string, bool) {
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
	if defaultCache != nil {
		defaultCache.Set("Shields"+key, r, cache.NoExpiration) // shields svg key相同，返回的数据一定相同
	}
}
