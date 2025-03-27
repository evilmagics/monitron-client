package stats

import (
	"github.com/shirou/gopsutil/v4/disk"
)

type DiskStat struct {
	Partitions []DiskPartition `json:"partitions"`
	Usage      DiskUsage       `json:"usage"`
}

type DiskUsage struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}

type DiskPartition struct {
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

func StatDisk() (DiskStat, error) {
	partitions, err := StatDiskPartitions()
	if err != nil {
		return DiskStat{}, err
	}

	usage, err := StatDiskUsage()
	if err != nil {
		return DiskStat{}, err
	}

	return DiskStat{
		Partitions: partitions,
		Usage:      usage,
	}, nil
}

func StatDiskUsage() (DiskUsage, error) {
	u, err := disk.Usage("/")
	if err != nil || u == nil {
		return DiskUsage{}, err
	}

	return DiskUsage{
		Total:       u.Total,
		Used:        u.Used,
		Free:        u.Free,
		UsedPercent: u.UsedPercent,
	}, nil
}

func StatDiskPartitions() ([]DiskPartition, error) {
	statPartitions := []DiskPartition{}

	partitions, err := disk.Partitions(false)
	if err != nil {
		return []DiskPartition{}, err
	}

	for _, p := range partitions {
		u, err := StatDiskPartition(p.Mountpoint, p.Device, p.Fstype)
		if err != nil {
			continue
		}

		statPartitions = append(statPartitions, u)
	}

	return statPartitions, nil
}
func StatDiskPartition(mountpoint string, opts ...string) (DiskPartition, error) {
	// Find partition information
	if len(opts) <= 2 {
		p := FindDiskPartition(mountpoint)
		opts = []string{p.Device, p.Fstype}
	}

	u, err := disk.Usage(mountpoint)
	if err != nil || u == nil {
		return DiskPartition{}, err
	}

	return DiskPartition{
		Mountpoint:        mountpoint,
		Device:            opts[0],
		Fstype:            opts[1],
		Total:             u.Total,
		Used:              u.Used,
		Free:              u.Free,
		UsedPercent:       u.UsedPercent,
		InodesTotal:       u.InodesTotal,
		InodesFree:        u.InodesFree,
		InodesUsed:        u.InodesUsed,
		InodesUsedPercent: u.InodesUsedPercent,
	}, err
}

func FindDiskPartition(mountpoint string) disk.PartitionStat {
	// Find partition information
	partitions, err := disk.Partitions(false)
	if err != nil {
		return disk.PartitionStat{}
	}

	for _, partition := range partitions {
		if partition.Mountpoint == mountpoint {
			return partition
		}
	}

	return disk.PartitionStat{}
}
