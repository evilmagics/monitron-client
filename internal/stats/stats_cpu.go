package stats

import (
	"github.com/goccy/go-json"
	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUInfo struct {
	VendorID      string  `json:"vendorId"`
	Family        string  `json:"family"`
	Model         string  `json:"model"`
	PhysicalID    string  `json:"physicalId"`
	Cores         int32   `json:"cores"`
	ModelName     string  `json:"modelName"`
	FrequencyUnit string  `json:"frequencyUnit"`
	Frequency     float64 `json:"frequency"`
	CacheSize     int32   `json:"cacheSize"`
}

func (s CPUInfo) String() string {
	str, _ := json.Marshal(s)
	return string(str)
}

type CPUUsage struct {
	User    float64 `json:"user"`
	Idle    float64 `json:"idle"`
	IOWait  float64 `json:"ioWait"`
	System  float64 `json:"system"`
	Nice    float64 `json:"nice"`
	Irq     float64 `json:"irq"`
	SoftIrq float64 `json:"softIrq"`
	Steal   float64 `json:"steal"`
	Guest   float64 `json:"guest"`
	TNice   float64 `json:"tNice"`
	Total   float64 `json:"total"`
	Percent float64 `json:"percent"`
}

func (s CPUUsage) String() string {
	str, _ := json.Marshal(s)
	return string(str)
}

type CPUStat struct {
	Info  CPUInfo  `json:"info"`
	Usage CPUUsage `json:"usage"`
}

func (s CPUStat) String() string {
	str, _ := json.Marshal(s)
	return string(str)
}

func StatCPU() (CPUStat, error) {
	info, err := cpu.Info()
	if err != nil || len(info) == 0 {
		return CPUStat{}, err
	}

	usage, err := cpu.Times(false)
	if err != nil || len(usage) == 0 {
		return CPUStat{}, err
	}

	percent, err := cpu.Percent(0, false)
	if err != nil || len(percent) == 0 {
		return CPUStat{}, err
	}

	stats := CPUStat{
		Info: CPUInfo{
			VendorID:      info[0].VendorID,
			Family:        info[0].Family,
			Model:         info[0].Model,
			PhysicalID:    info[0].PhysicalID,
			ModelName:     info[0].ModelName,
			Cores:         info[0].Cores,
			Frequency:     info[0].Mhz / 1000.0,
			CacheSize:     info[0].CacheSize,
			FrequencyUnit: "GHz",
		},
		Usage: CPUUsage{
			User:    usage[0].User,
			System:  usage[0].System,
			Idle:    usage[0].Idle,
			IOWait:  usage[0].Iowait,
			Nice:    usage[0].Nice,
			Irq:     usage[0].Irq,
			SoftIrq: usage[0].Softirq,
			Steal:   usage[0].Steal,
			Guest:   usage[0].Guest,
			TNice:   usage[0].Nice,
			Total:   usage[0].Total(),
			Percent: percent[0],
		},
	}
	return stats, nil
}
