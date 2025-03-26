package database

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var once sync.Once
var memcache *cache.Cache

func Cache() *cache.Cache {
	if memcache == nil {
		once.Do(func() {
			memcache = cache.New(1*time.Minute, 3*time.Minute)
		})
	}
	return memcache
}
