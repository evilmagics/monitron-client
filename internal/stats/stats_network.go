package stats

import (
	"github.com/shirou/gopsutil/v4/net"
)

type NetworkStat struct {
	Usage NetworkUsage `json:"usage"`
}
type NetworkUsage struct {
	BytesSent       uint64 `json:"bytesSent"`
	BytesReceived   uint64 `json:"bytesReceived"`
	PacketsSent     uint64 `json:"packetsSent"`
	PacketsReceived uint64 `json:"packetsReceived"`
	ErrorIn         uint64 `json:"errorIn"`
	ErrorOut        uint64 `json:"errorOut"`
	DropIn          uint64 `json:"dropIn"`
	DropOut         uint64 `json:"dropOut"`
	FifoIn          uint64 `json:"fifoIn"`
	FifoOut         uint64 `json:"fifoOut"`
}

func StatNetwork() (NetworkStat, error) {
	usage, err := StatNetworkUsage()
	if err != nil {
		return NetworkStat{}, err
	}

	return NetworkStat{Usage: usage}, nil
}

func StatNetworkUsage() (NetworkUsage, error) {
	counters, err := net.IOCounters(false)
	if err != nil || len(counters) == 0 {
		return NetworkUsage{}, err
	}

	return NetworkUsage{
		BytesSent:       counters[0].BytesSent,
		BytesReceived:   counters[0].BytesRecv,
		PacketsSent:     counters[0].PacketsSent,
		PacketsReceived: counters[0].PacketsRecv,
		ErrorIn:         counters[0].Errin,
		ErrorOut:        counters[0].Errout,
		DropIn:          counters[0].Dropin,
		DropOut:         counters[0].Dropout,
		FifoIn:          counters[0].Fifoin,
		FifoOut:         counters[0].Fifoout,
	}, nil
}
