package windowsagent

import (
	"encoding/json"
	"net"

	"github.com/Expand-My-Business/go_windows_agent/utils"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	next "github.com/shirou/gopsutil/net"
	"github.com/sirupsen/logrus"
)

type Addr struct {
	MacAddr   string `json:"mac_address"`
	PublicIp  string `json:"public_ip"`
	PrivateIp string `json:"private_ip"`
}

type Stats struct {
	CpuInfo     []cpu.InfoStat         `json:"cpuInfo,omitempty"`
	DiskInfo    *disk.UsageStat        `json:"diskInfo,omitempty"`
	HostInfo    *host.InfoStat         `json:"hostInfo,omitempty"`
	MemoryInfo  *mem.VirtualMemoryStat `json:"memoryInfo,omitempty"`
	NetworkInfo []next.IOCountersStat  `json:"networkInfo,omitempty"`
	Addr        Addr                   `json:"address"`
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
	macAddr, err := utils.GetMacAddresses()
	if err != nil {
		logrus.Errorf("cannot get the mac address")
	}

	pubAddr, err := utils.GetPublicIP()
	if err != nil {
		logrus.Errorf("cannot get the mac address")
	}

	privAddr, err := utils.GetPrivateIPAddress()
	if err != nil {
		logrus.Errorf("cannot get the mac address")
	}
	addr := Addr{
		MacAddr:   macAddr,
		PublicIp:  pubAddr,
		PrivateIp: privAddr,
	}
	windowsStats.CpuInfo = cpu
	windowsStats.DiskInfo = disk
	windowsStats.HostInfo = host
	windowsStats.MemoryInfo = memory
	windowsStats.NetworkInfo = network
	windowsStats.HostIP = out
	windowsStats.Addr = addr

	windowsByteSlice, err := json.MarshalIndent(windowsStats, "", "\t")
	if err != nil {
		logrus.Errorf("cannot marshal to byteslice", err)
		return nil, err
	}

	// if err := ioutil.WriteFile("address.json", windowsByteSlice, 0777); err != nil {
	// 	logrus.Errorf("cannot write address.json, error: %+v", err)
	// }
	return windowsByteSlice, nil

}

func CpuStats() ([]cpu.InfoStat, error) {
	info, err := cpu.Info()

	//almost every return value is a struct
	// fmt.Printf("All info: ", info)
	return info, err
}
