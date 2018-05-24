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

//func cacheGet(key string) {
//	r, err := defaultCache.Get(key)
//	if err != nil {
//		return
//	}
//
//}
