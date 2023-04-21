package netstat

import (
	"encoding/json"
	"fmt"

	"github.com/Expand-My-Business/go_windows_agent/netstat/commands"
	"github.com/Expand-My-Business/go_windows_agent/utils"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"github.com/sirupsen/logrus"
)

type NetStatsDetails struct {
	NetStats NetStats `json:"sys_ports"`
	HostIP   string   `json:"hostIP"`
}

type NetStats struct {
	UDPStats  []UDPStats
	TCPStats  []TCPStats
	BGProcess []commands.Process
}

type TCPStats struct {
	Name     string   `json:"processname"`
	Family   uint32   `json:"family"`
	Type     uint32   `json:"type"`
	Laddr    net.Addr `json:"localaddr"`
	Raddr    net.Addr `json:"remoteaddr"`
	Status   string   `json:"status"`
	Uids     []int32  `json:"uids"`
	Pid      int32    `json:"pid"`
	Protocol string   `json:"protocol"`
}

type UDPStats struct {
	Name     string   `json:"processname"`
	Family   uint32   `json:"family"`
	Type     uint32   `json:"type"`
	Laddr    net.Addr `json:"localaddr"`
	Raddr    net.Addr `json:"remoteaddr"`
	Status   string   `json:"status"`
	Uids     []int32  `json:"uids"`
	Pid      int32    `json:"pid"`
	Protocol string   `json:"protocol"`
}

func GetNetStats() ([]byte, error) {
	// Get TCP connections
	tcpConns, err := net.Connections("tcp")
	if err != nil {
		panic(err)
	}

	allTCPStats := []TCPStats{}
	// Print TCP connections with process names
	fmt.Println("Getting stats for TCP connections")
	for _, conn := range tcpConns {
		proc, err := process.NewProcess(conn.Pid)
		if err != nil {
			panic(err)
		}
		name, err := proc.Name()
		if err != nil {
			panic(err)
		}
		tcpStats := TCPStats{}
		tcpStats.Name = name
		tcpStats.Family = conn.Family
		tcpStats.Type = conn.Type
		tcpStats.Laddr.IP = conn.Laddr.IP
		tcpStats.Laddr.Port = conn.Laddr.Port
		tcpStats.Raddr.IP = conn.Raddr.IP
		tcpStats.Raddr.Port = conn.Raddr.Port
		tcpStats.Status = conn.Status
		tcpStats.Pid = conn.Pid
		tcpStats.Protocol = "TCP"
		allTCPStats = append(allTCPStats, tcpStats)
	}

	// Get UDP connections
	udpConns, err := net.Connections("udp")
	if err != nil {
		panic(err)
	}

	allUDPStats := []UDPStats{}

	// Print UDP connections with process names
	fmt.Println("Getting stats for UDP connections")
	for _, conn := range udpConns {
		proc, err := process.NewProcess(conn.Pid)
		if err != nil {
			panic(err)
		}
		name, err := proc.Name()
		if err != nil {
			panic(err)
		}

		udpStats := UDPStats{}
		udpStats.Name = name
		udpStats.Family = conn.Family
		udpStats.Type = conn.Type
		udpStats.Laddr.IP = conn.Laddr.IP
		udpStats.Laddr.Port = conn.Laddr.Port
		udpStats.Raddr.IP = conn.Raddr.IP
		udpStats.Raddr.Port = conn.Raddr.Port
		udpStats.Status = conn.Status
		udpStats.Pid = conn.Pid
		udpStats.Protocol = "UDP"
		allUDPStats = append(allUDPStats, udpStats)
	}

	processes, err := commands.GetAllInternalProcess()
	if err != nil {
		logrus.Errorf("cannot get background process, error: %s", err)
	}

	netStatDetails := NetStatsDetails{}
	netStatDetails.NetStats = NetStats{
		UDPStats:  allUDPStats,
		TCPStats:  allTCPStats,
		BGProcess: processes,
	}

	// Get private Ip
	ip, err := utils.GetPrivateIPAddress()
	if err != nil {
		logrus.Errorf("cannot get private ip")
	}

	netStatDetails.HostIP = ip

	byteSlice, err := json.MarshalIndent(netStatDetails, "", "\t")
	if err != nil {
		logrus.Errorf("cannot marshal net stat details: %+v", err)
		return nil, err
	}

	return byteSlice, err
}
