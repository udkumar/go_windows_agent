package main

import (
	"fmt"

	"github.com/go_windows_agent/windows_agent"
)

func main() {
	// winServices, _ := windowsagent.WindowsServices()

	cpuInfo, _ := windowsagent.CpuStats()
	fmt.Println("cpu info", cpuInfo)
	
	
} 
