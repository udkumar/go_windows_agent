package main

import (
	"bytes"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Expand-My-Business/go_windows_agent/netstat"
	"github.com/Expand-My-Business/go_windows_agent/nmap"
	"github.com/Expand-My-Business/go_windows_agent/utils"
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
	currentWD, err := utils.GetWorkingDir()
	if err != nil {
		logrus.Errorf("cannot get the current working dir, error: %+v", err)
	}

	logfilePath := filepath.Join(currentWD, "agent.log")
	// Open a file for writing store the log file in current working directory
	file, err := os.OpenFile(logfilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
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

	// ioutil.WriteFile("nmapXbyte.json", nmapXbyte, 0777)
	// ioutil.WriteFile("winXbytes.json", winXbytes, 0777)
	// ioutil.WriteFile("netXbyte.json", netXbyte, 0777)

	// Send json value to certain API and certain interval
	for {

		go sendStringToAPI("http://13.235.66.99/agent_ports_data", string(nmapXbyte))
		go sendStringToAPI("http://13.235.66.99/add_agent_logs", string(winXbytes))
		go sendStringToAPI("http://13.235.66.99/agent_process_data", string(netXbyte))

		time.Sleep(time.Minute * 3)
	}

}

func sendStringToAPI(url string, data string) error {
	logrus.Infof("Sending data to API: %s", url)
	requestBody := bytes.NewBuffer([]byte(data))

	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		logrus.Errorf("cannot make a request wrapper, error: %+v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("cannot send a request, error: %+v", err)
		return err
	}
	defer resp.Body.Close()
	logrus.Infof("Ending execution for API: %s", url)
	return nil
}
