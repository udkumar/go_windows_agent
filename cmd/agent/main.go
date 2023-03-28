package main

import (
	"bytes"
	"net"
	"net/http"
	"time"

	"github.com/Expand-My-Business/go_windows_agent/nmap_stack"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	next "github.com/shirou/gopsutil/net"
)

type Stats struct {
	CpuInfo     []cpu.InfoStat         `json:"cpuInfo,omitempty"`
	DiskInfo    *disk.UsageStat        `json:"diskInfo,omitempty"`
	HostInfo    *host.InfoStat         `json:"hostInfo,omitempty"`
	MemoryInfo  *mem.VirtualMemoryStat `json:"memoryInfo,omitempty"`
	NetworkInfo []next.IOCountersStat  `json:"networkInfo,omitempty"`
	HostIP      net.IP                 `json:"hostIP,omitempty"`
}

func main() {
	jsonVal := nmap_stack.Stats("127.0.0.1", "1-1000")

	_ = jsonVal

	// Send json value to certain API and certain interval
	for {
		// TODO: change empty string to api endpoint
		sendStringToAPI("", jsonVal)

		time.Sleep(time.Minute * 10)
	}

}

func sendStringToAPI(url string, data string) error {
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

	return nil
}
