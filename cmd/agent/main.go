package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"mime"
	"net/http"

	windowsagent "github.com/udkumar/go_windows_agent/windows_agent"
)

func main() {
	// winServices, _ := windowsagent.WindowsServices()
	flag.Parse()
	reqURL := flag.Arg(0)

	cpuInfo, _ := windowsagent.CpuStats()
	fmt.Println("cpu info", cpuInfo)

	if reqURL != "" {
		body := map[string]interface{}{
			"cpuInfo": cpuInfo,
		}

		buf := &bytes.Buffer{}
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			fmt.Println("Failed to encode", err)
			return
		}

		resp, err := http.Post(reqURL, mime.TypeByExtension(".json"), buf)
		if err != nil {
			fmt.Println("Failed to request", err)
		} else {
			fmt.Println("Response", resp.StatusCode)
		}
	}
}
