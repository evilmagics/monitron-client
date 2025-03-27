package stats

import (
	"sync"
	"time"
)

var once sync.Once
var ticker *time.Ticker

// Caching stats each 'n' interval on memory until 'm' durations
func StartCollectStats(interval ...time.Duration) {
	exp := DefaultTickerInterval + time.Second
	if len(interval) > 0 {
		exp = interval[0] + time.Second
		SetTickerInterval(interval[0])
	}

	// Cache first
	CacheUsageStats(exp)

	for range Ticker().C {
		CacheUsageStats(exp)
	}
}

func StopCollectStats() {
	Ticker().Stop()
}

func SetTickerInterval(interval time.Duration) {
	Ticker().Reset(interval)
}

func Ticker() *time.Ticker {
	if ticker == nil {
		once.Do(func() {
			ticker = time.NewTicker(DefaultTickerInterval)
		})
	}
	return ticker
}
