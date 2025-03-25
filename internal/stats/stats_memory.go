package stats

import (
	"github.com/goccy/go-json"
	"github.com/shirou/gopsutil/v4/mem"
)

type MemoryStat struct {
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	Cached      uint64  `json:"cached"`
	SwapFree    uint64  `json:"swapFree"`
	SwapCached  uint64  `json:"swapCached"`
	SwapTotal   uint64  `json:"swapTotal"`
	UsedPercent float64 `json:"usedPercent"`
}

func (s MemoryStat) String() string {
	str, _ := json.Marshal(s)
	return string(str)
}

func StatMemory() (MemoryStat, error) {
	v, err := mem.VirtualMemory()
	if err != nil || v == nil {
		return MemoryStat{}, err
	}

	return MemoryStat{
		Available:   v.Available,
		Used:        v.Used,
		Free:        v.Free,
		Cached:      v.Cached,
		SwapFree:    v.SwapFree,
		SwapCached:  v.SwapCached,
		SwapTotal:   v.SwapTotal,
		UsedPercent: v.UsedPercent,
	}, nil
}
