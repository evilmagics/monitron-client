package stats

import (
	"github.com/goccy/go-json"
)

type Stats struct {
	Host    HostStat    `json:"host"`
	CPU     CPUStat     `json:"cpu"`
	Memory  MemoryStat  `json:"memory"`
	Disk    DiskStat    `json:"disk"`
	Network NetworkStat `json:"network"`
}

func (s Stats) Json() ([]byte, error) {
	return json.Marshal(s)
}
func (s Stats) String() string {
	str, _ := s.Json()
	return string(str)
}

func AllStats() (*Stats, error) {
	var err error
	stats := new(Stats)

	stats.CPU, err = StatCPU()
	stats.Memory, err = StatMemory()
	stats.Disk, err = StatDisk()
	stats.Network, err = StatNetwork()

	return stats, err
}
