package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/Expand-My-Business/go_windows_agent/nmap_stack"
	"github.com/Expand-My-Business/go_windows_agent/windowsagent"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	next "github.com/shirou/gopsutil/net"
)

type AllStats struct {
	WinStats  Stats
	NmapStats nmap_stack.NmapStats
}

type Stats struct {
	CpuInfo     []cpu.InfoStat         `json:"cpuInfo,omitempty"`
	DiskInfo    *disk.UsageStat        `json:"diskInfo,omitempty"`
	HostInfo    *host.InfoStat         `json:"hostInfo,omitempty"`
	MemoryInfo  *mem.VirtualMemoryStat `json:"memoryInfo,omitempty"`
	NetworkInfo []next.IOCountersStat  `json:"networkInfo,omitempty"`
	HostIP      net.IP                 `json:"hostIP,omitempty"`
}

// func main() {
// 	flag.Parse()
// 	reqURL := flag.Arg(0)
// 	fmt.Printf("*****reqURL: %v\n", reqURL)

// 	for {
// 		fmt.Println("https://gosamples.dev is the best")

// 		// stats := &Stats{}
// 		// cpu, _ := windowsagent.CpuStats()
// 		// disk, _ := windowsagent.DiskStats()
// 		// host, _ := windowsagent.HostStats()
// 		// memory, _ := windowsagent.MemoryStats()
// 		// network, _ := windowsagent.NetStats()
// 		// out := windowsagent.GetOutboundIP()

// 		// stats.CpuInfo = cpu
// 		// stats.DiskInfo = disk
// 		// stats.HostInfo = host
// 		// stats.MemoryInfo = memory
// 		// stats.NetworkInfo = network
// 		// stats.HostIP = out
// 		// fmt.Printf("*****cpu: %v\n", cpu)

// 		jsonVal := nmap_stack.Stats("127.0.0.1", "1-1000")
// 		fmt.Printf("jsonVal: %v\n", jsonVal)
// 		_ = jsonVal

// 		allStats := AllStats{
// 			NmapStats: jsonVal,
// 		}

// 		go sendStringToAPI("http://13.235.66.99/agent_ports_data", jsonVal)

// 		// bx, err := json.MarshalIndent(stats, "", "    ")
// 		// if err != nil {
// 		// 	fmt.Errorf("cannot marshal to byteslice", err)
// 		// }

// 		// fmt.Println("##.....", string(bx))

// 		// ioutil.WriteFile("demo.json", bx, 0777)

// 		// if reqURL == "" {

// 		// 	params := url.Values{}
// 		// 	params.Add("Agent Info", string(bx))

// 		// 	resp, _ := http.PostForm("https://f232-2401-4900-8097-8a91-6551-8460-6de4-248a.in.ngrok.io/api/saveData", params)
// 		// 	if err != nil {
// 		// 		fmt.Println("Failed to request", err)
// 		// 		return
// 		// 	} else {
// 		// 		fmt.Println("Response", resp)
// 		// 	}

// 		// }

// 		// time.Sleep(2 * time.Second)
// 	}
// }

func main() {
	jsonVal := nmap_stack.Stats("127.0.0.1", "1-1000")

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

	bxStats, err := json.MarshalIndent(stats, "", "    ")
	if err != nil {
		fmt.Errorf("cannot marshal to byteslice", err)
	}
	// Send json value to certain API and certain interval
	for {

		go sendStringToAPI("http://13.235.66.99/agent_ports_data", jsonVal)
		go sendStringToAPI("http://13.235.66.99/add_agent_logs", string(bxStats))

		time.Sleep(time.Minute * 3)
	}

}

func sendStringToAPI(url string, data string) error {
	fmt.Printf("Sending data to API: ", url)
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
