package stats

import (
	"github.com/shirou/gopsutil/v4/host"
)

type HostStat struct {
	HostID          string `json:"hostId"`
	Hostname        string `json:"hostname"`
	Uptime          uint64 `json:"uptime"`
	BootTime        uint64 `json:"bootTime"`
	Procs           uint64 `json:"procs"`
	OS              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformFamily  string `json:"platformFamily"`
	PlatformVersion string `json:"platformVersion"`
	KernelVersion   string `json:"kernelVersion"`
	KernelArch      string `json:"kernelArch"`
}

func StatHost() (HostStat, error) {
	info, err := host.Info()
	if err != nil || info == nil {
		return HostStat{}, err
	}

	return HostStat{
		HostID:          info.HostID,
		Hostname:        info.Hostname,
		Uptime:          info.Uptime,
		BootTime:        info.BootTime,
		Procs:           info.Procs,
		OS:              info.OS,
		Platform:        info.Platform,
		PlatformFamily:  info.PlatformFamily,
		PlatformVersion: info.PlatformVersion,
		KernelVersion:   info.KernelVersion,
		KernelArch:      info.KernelArch,
	}, nil
}
