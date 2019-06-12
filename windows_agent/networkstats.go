package windowsagent

import (
    "fmt"

    "github.com/shirou/gopsutil/net"
)

func NetStats() {
    infoIOCounter, _ := net.IOCounters(true)

    //almost every return value is a struct
    fmt.Printf("All net info: ", infoIOCounter)
    return infoIOCounter
    
}