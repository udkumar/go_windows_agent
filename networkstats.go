package net

import (
    "fmt"

    "github.com/shirou/gopsutil/net"
)

func main() {
    infoIOCounter, _ := net.IOCounters(true)

    //almost every return value is a struct
    fmt.Printf("All net info: ", infoIOCounter)
    
}