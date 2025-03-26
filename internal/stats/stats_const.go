package stats

import "time"

type StatFunc[T any] func() (T, error)
type GetCacheFunc[T any] func() *T

const (
	DefaultTickerInterval = time.Second
	DefaultCacheExp       = time.Second
)

const (
	CacheKeyStats   = "stats"
	CacheKeyMemory  = "stats_memory"
	CacheKeyCPU     = "stats_cpu"
	CacheKeyDisk    = "stats_disk"
	CacheKeyNetwork = "stats_network"
)
