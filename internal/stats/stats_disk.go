package stats

import (
	"github.com/shirou/gopsutil/v4/disk"
)

type DiskUsage struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}
type DiskUsagePartition struct {
	Device            string  `json:"device"`
	Mountpoint        string  `json:"mountpoint"`
	Fstype            string  `json:"fstype"`
	Total             uint64  `json:"total"`
	Used              uint64  `json:"used"`
	Free              uint64  `json:"free"`
	UsedPercent       float64 `json:"usedPercent"`
	InodesTotal       uint64  `json:"inodesTotal"`
	InodesFree        uint64  `json:"inodesFree"`
	InodesUsed        uint64  `json:"inodesUsed"`
	InodesUsedPercent float64 `json:"inodesUsedPercent"`
}
type DiskStat struct {
	Partitions []DiskUsagePartition `json:"partitions"`
	Usage      DiskUsage            `json:"usage"`
}

func StatDisk() (DiskStat, error) {
	usage, err := disk.Usage("/")
	if err != nil || usage == nil {
		return DiskStat{}, err
	}

	partitions, err := disk.Partitions(true)
	if err != nil {
		return DiskStat{}, err
	}

	partitionUsage := []DiskUsagePartition{}

	for _, p := range partitions {
		u, err := disk.Usage(p.Mountpoint)
		if err != nil || u == nil {
			continue
		}
		partitionUsage = append(partitionUsage, DiskUsagePartition{
			Device:            p.Device,
			Mountpoint:        p.Mountpoint,
			Fstype:            p.Fstype,
			Total:             u.Total,
			Used:              u.Used,
			Free:              u.Free,
			UsedPercent:       u.UsedPercent,
			InodesTotal:       u.InodesTotal,
			InodesFree:        u.InodesFree,
			InodesUsed:        u.InodesUsed,
			InodesUsedPercent: u.InodesUsedPercent,
		})
	}

	return DiskStat{
		Partitions: partitionUsage,
		Usage: DiskUsage{
			Total:       usage.Total,
			Used:        usage.Used,
			Free:        usage.Free,
			UsedPercent: usage.UsedPercent,
		},
	}, nil
}
