package windowsagent

import (
	"encoding/json"
	"net"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	next "github.com/shirou/gopsutil/net"
	"github.com/sirupsen/logrus"
)

type Stats struct {
	CpuInfo     []cpu.InfoStat         `json:"cpuInfo,omitempty"`
	DiskInfo    *disk.UsageStat        `json:"diskInfo,omitempty"`
	HostInfo    *host.InfoStat         `json:"hostInfo,omitempty"`
	MemoryInfo  *mem.VirtualMemoryStat `json:"memoryInfo,omitempty"`
	NetworkInfo []next.IOCountersStat  `json:"networkInfo,omitempty"`
	HostIP      net.IP                 `json:"hostIP,omitempty"`
}

func GetWindowsStats() ([]byte, error) {
	windowsStats := &Stats{}
	cpu, err := CpuStats()
	if err != nil {
		logrus.Errorf("cannot get CPU stats, error: %+v", err)
	}
	disk, err := DiskStats()
	if err != nil {
		logrus.Errorf("cannot get Disk stats, error: %+v", err)
	}

	host, err := HostStats()
	if err != nil {
		logrus.Errorf("cannot get Host stats, error: %+v", err)
	}

	memory, err := MemoryStats()
	if err != nil {
		logrus.Errorf("cannot get Memory stats, error: %+v", err)
	}

	network, err := NetStats()
	if err != nil {
		logrus.Errorf("cannot get Net stats, error: %+v", err)
	}

	out := GetOutboundIP()

	windowsStats.CpuInfo = cpu
	windowsStats.DiskInfo = disk
	windowsStats.HostInfo = host
	windowsStats.MemoryInfo = memory
	windowsStats.NetworkInfo = network
	windowsStats.HostIP = out

	windowsByteSlice, err := json.MarshalIndent(windowsStats, "", "\t")
	if err != nil {
		logrus.Errorf("cannot marshal to byteslice", err)
		return nil, err
	}

	return windowsByteSlice, nil

}

func CpuStats() ([]cpu.InfoStat, error) {
	info, err := cpu.Info()

	//almost every return value is a struct
	// fmt.Printf("All info: ", info)
	return info, err
}
