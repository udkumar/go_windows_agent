package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	next "github.com/shirou/gopsutil/net"
	windowsagent "github.com/udkumar/go_windows_agent/windows_agent"
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
	flag.Parse()
	reqURL := flag.Arg(0)
	fmt.Printf("*****reqURL: %v\n", reqURL)

	for {
		fmt.Println("https://gosamples.dev is the best")

		stats := &Stats{}
		cpu, _ := windowsagent.CpuStats()
		disk, _ := windowsagent.DiskStats()
		host, _ := windowsagent.HostStats()
		memory, _ := windowsagent.MemoryStats()
		network, _ := windowsagent.NetStats()
		out := windowsagent.GetOutboundIP()

		stats.CpuInfo = cpu
		stats.DiskInfo = disk
		stats.HostInfo = host
		stats.MemoryInfo = memory
		stats.NetworkInfo = network
		stats.HostIP = out
		fmt.Printf("*****cpu: %v\n", cpu)

		bx, err := json.MarshalIndent(stats, "", "    ")
		if err != nil {
			fmt.Errorf("cannot marshal to byteslice", err)
		}

		fmt.Println("##.....", string(bx))

		ioutil.WriteFile("demo.json", bx, 0777)

		if reqURL == "" {

			params := url.Values{}
			params.Add("Agent Info", string(bx))

			resp, _ := http.PostForm("https://f232-2401-4900-8097-8a91-6551-8460-6de4-248a.in.ngrok.io/api/saveData", params)
			if err != nil {
				fmt.Println("Failed to request", err)
				return
			} else {
				fmt.Println("Response", resp)
			}

		}

		time.Sleep(2 * time.Second)
	}
}
