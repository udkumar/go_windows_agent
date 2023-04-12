package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/Expand-My-Business/go_windows_agent/netstat"
	"github.com/Expand-My-Business/go_windows_agent/nmap"
	"github.com/Expand-My-Business/go_windows_agent/windowsagent"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	next "github.com/shirou/gopsutil/net"
	"github.com/sirupsen/logrus"
)

type AllStats struct {
	WinStats  Stats
	NmapStats nmap.NmapStats
}

type Stats struct {
	CpuInfo     []cpu.InfoStat         `json:"cpuInfo,omitempty"`
	DiskInfo    *disk.UsageStat        `json:"diskInfo,omitempty"`
	HostInfo    *host.InfoStat         `json:"hostInfo,omitempty"`
	MemoryInfo  *mem.VirtualMemoryStat `json:"memoryInfo,omitempty"`
	NetworkInfo []next.IOCountersStat  `json:"networkInfo,omitempty"`
	HostIP      net.IP                 `json:"hostIP,omitempty"`
}

func main() {
	// Open a file for writing
	file, err := os.OpenFile("agent.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Errorf("cannot open the logfile, error: %+v", err)
	}
	defer file.Close()

	// Set the log output to the file
	logrus.SetOutput(file)

	nmapXbyte, err := nmap.GetNmapDetails("127.0.0.1", "1-1000")
	if err != nil {
		logrus.Errorf("cannot get nmap details, error: %+v", err)
	}

	winXbytes, err := windowsagent.GetWindowsStats()
	if err != nil {
		logrus.Errorf("cannot get windows stats, error: %+v", nil)
	}
	netXbyte, err := netstat.GetNetStats()
	if err != nil {
		logrus.Errorf("cannot get netstat details, error: %+v", nil)
	}
	// windowsStats := &Stats{}
	// cpu, _ := windowsagent.CpuStats()
	// disk, _ := windowsagent.DiskStats()
	// host, _ := windowsagent.HostStats()
	// memory, _ := windowsagent.MemoryStats()
	// network, _ := windowsagent.NetStats()
	// out := windowsagent.GetOutboundIP()

	// windowsStats.CpuInfo = cpu
	// windowsStats.DiskInfo = disk
	// windowsStats.HostInfo = host
	// windowsStats.MemoryInfo = memory
	// windowsStats.NetworkInfo = network
	// windowsStats.HostIP = out

	// bxStats, err := json.MarshalIndent(windowsStats, "", "    ")
	// if err != nil {
	// 	logrus.Errorf("cannot marshal to byteslice", err)
	// }

	// netStats, err := netstat.Netstat()
	// if err != nil {
	// 	logrus.Errorf("cannot marshal to byteslice", err)
	// }

	// nsBytes, err := json.MarshalIndent(netStats, "", "    ")
	// if err != nil {
	// 	logrus.Errorf("cannot marshal to byteslice", err)
	// }

	ioutil.WriteFile("nmapXbyte.json", nmapXbyte, 0777)
	ioutil.WriteFile("winXbytes.json", winXbytes, 0777)
	ioutil.WriteFile("netXbyte.json", netXbyte, 0777)
	// ioutil.WriteFile("netStats.json", nsBytes, 0777)
	// Send json value to certain API and certain interval
	for {

		go sendStringToAPI("http://13.235.66.99/agent_ports_data", string(nmapXbyte))
		go sendStringToAPI("http://13.235.66.99/add_agent_logs", string(winXbytes))
		go sendStringToAPI("http://13.235.66.99/agent_process_data", string(netXbyte))

		time.Sleep(time.Minute * 3)
	}

}

func sendStringToAPI(url string, data string) error {
	logrus.Infof("Sending data to API: ", url)
	requestBody := bytes.NewBuffer([]byte(data))

	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Printf("Ending execution for API: ", url)
	return nil
}
