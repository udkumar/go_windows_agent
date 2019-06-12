package windowsagent

import (
    "fmt"

    "github.com/shirou/gopsutil/net"
)

func NetStats() ([]net.IOCountersStat, error) {
    infoIOCounter, err := net.IOCounters(true)

    //almost every return value is a struct
    fmt.Printf("All net info: ", infoIOCounter)
    return infoIOCounter, err
    
}