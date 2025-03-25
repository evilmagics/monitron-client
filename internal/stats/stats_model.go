package stats

type StatsModel struct {
	MemoryAvailable string `json:"memory_available"`
}

// type MemoryStat struct {
// 	Total       int64   `json:"total"`
// 	Free        int64   `json:"free"`
// 	Used        int64   `json:"used"`
// 	UsedPercent float64 `json:"usedPercent"`
// 	Cached      float64 `json:"cached"`
// }

// type DiskStat struct {
// 	Path        string  `json:"path"`
// 	Total       int64   `json:"total"`
// 	Free        int64   `json:"free"`
// 	Used        int64   `json:"used"`
// 	UsedPercent float32 `json:"usedPercent"`
// }

// type HostStat struct {
// 	ID              string `json:"id"`
// 	Hostname        string `json:"hostname"`
// 	Uptime          int64  `json:"uptime"`
// 	BootTime        int64  `json:"bootTime"`
// 	Procs           int    `json:"procs"`
// 	OS              string `json:"os"`
// 	Platform        string `json:"platform"`
// 	PlatformVersion string `json:"platformVersion"`
// 	KernelVersion   string `json:"kernelVersion"`
// 	KernelArch      string `json:"kernelArch"`
// }
