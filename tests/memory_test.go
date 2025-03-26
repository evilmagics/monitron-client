package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

func TestInMemoryCache(t *testing.T) {
	c := cache.New(1*time.Minute, 3*time.Minute)

	// stat := stats.CollectStats()
	stat := -1

	setTime := time.Now()
	// Test adding and retrieving items from the cache
	c.Set("stats", stat, time.Second)
	ss, _ := c.Get("stats")
	fmt.Println(ss)

	stat = 2
	c.Set("stats", stat, time.Second)
	ss, _ = c.Get("stats")
	fmt.Println(ss)

	fmt.Printf("\nSet time: %s\n", time.Since(setTime))

	wg := new(sync.WaitGroup)
	retTime := time.Now()
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Retrieve items from cache
			if x, found := c.Get("stats"); found {
				s := x.(*int)
				fmt.Println(s)
			}
		}()

	}
	wg.Wait()
	fmt.Printf("Retrieve time: %s", time.Since(retTime))

	t.Fail()
}
