package tests

import (
	"sync"
	"testing"
	"time"
)

var once sync.Once
var ticker *time.Ticker

func Ticker() *time.Ticker {
	if ticker == nil {
		once.Do(func() { ticker = time.NewTicker(time.Second) })
	}
	return ticker
}

func TestTicker(t *testing.T) {
	for range Ticker().C {
		t.Log("Tick")
	}

	t.Fail()
}
