package main

import (
	"fmt"

	windowsagent "github.com/udkumar/go_windows_agent/windows_agent"
)

func main() {
	// winServices, _ := windowsagent.WindowsServices()

	cpuInfo, _ := windowsagent.CpuStats()
	fmt.Println("cpu info", cpuInfo)
}
