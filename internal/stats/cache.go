package stats

import (
	"monitron-client/internal/database"
	"time"
)

func GetCachedCPUUsage() *CPUUsage {
	return getCached[CPUUsage](CacheKeyCPU)

}
func GetCachedMemoryUsage() *MemoryUsage {
	return getCached[MemoryUsage](CacheKeyMemory)
}
func GetCachedDiskUsage() *DiskUsage {
	return getCached[DiskUsage](CacheKeyDisk)
}
func GetCachedNetworkUsage() *NetworkUsage {
	return getCached[NetworkUsage](CacheKeyNetwork)
}

func CacheUsageStats(exp ...time.Duration) {
	CacheCPUUsage(exp...)
	CacheMemoryUsage(exp...)
	CacheDiskUsage(exp...)
	CacheNetworkUsage(exp...)
}

func CacheCPUUsage(exp ...time.Duration) {
	caching(CacheKeyCPU, StatCPUUsage, exp...)
}

func CacheMemoryUsage(exp ...time.Duration) {
	caching(CacheKeyMemory, StatMemoryUsage, exp...)
}

func CacheDiskUsage(exp ...time.Duration) {
	caching(CacheKeyDisk, StatDiskUsage, exp...)
}

func CacheNetworkUsage(exp ...time.Duration) {
	caching(CacheKeyNetwork, StatNetworkUsage, exp...)
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
