package tests

import (
	"fmt"
	"monitron-client/internal/stats"
	"testing"
	"time"
)

func TestStatCollection(t *testing.T) {
	go func() {
		fmt.Println("Start collect stats")
		stats.StartCollectStats()
	}()

	// Print output
	go func() {
		fmt.Println("Start print out stats")
		for range stats.Ticker().C {
			s := stats.GetCachedMemory()
			fmt.Printf("Received: %+v\n", s)
		}
	}()

	time.Sleep(3 * time.Second)
	stats.SetTickerInterval(200 * time.Millisecond)

	time.Sleep(3 * time.Second)
	stats.StopCollectStats()
	t.Fail()

}
