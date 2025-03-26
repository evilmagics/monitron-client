package stats

import (
	"monitron-client/internal/database"
	"time"
)

func GetCachedAllStats() *Stats {
	return getCached[Stats](CacheKeyStats)
}
func GetCachedCPU() *CPUStat {
	return getCached[CPUStat](CacheKeyCPU)
}
func GetCachedMemory() *MemoryStat {
	return getCached[MemoryStat](CacheKeyMemory)
}
func GetCachedDisk() *DiskStat {
	return getCached[DiskStat](CacheKeyDisk)
}
func GetCachedNetwork() *NetworkStat {
	return getCached[NetworkStat](CacheKeyNetwork)
}

func CacheAllStats(exp ...time.Duration) {
	CacheCPU(exp...)
	CacheMemory(exp...)
	CacheDisk(exp...)
	CacheNetwork(exp...)
}

func CacheCPU(exp ...time.Duration) {
	caching(CacheKeyCPU, StatCPU, exp...)
}

func CacheMemory(exp ...time.Duration) {
	caching(CacheKeyMemory, StatMemory, exp...)
}

func CacheDisk(exp ...time.Duration) {
	caching(CacheKeyDisk, StatDisk, exp...)
}

func CacheNetwork(exp ...time.Duration) {
	caching(CacheKeyNetwork, StatNetwork, exp...)
}

func caching[T any](key string, statFn StatFunc[T], exp ...time.Duration) {
	if len(exp) == 0 {
		exp = []time.Duration{DefaultCacheExp}
	}

	if stat, err := statFn(); err == nil {
		database.Cache().Set(key, &stat, exp[0])
	}

}

func getCached[T any](key string) *T {
	if s, ok := database.Cache().Get(key); ok && s != nil {
		return s.(*T)
	}
	return nil
}
