package main

import (
    "fmt"

    "github.com/shirou/gopsutil/host"
)

func main() {
    infoStat, _ := host.Info()

    //almost every return value is a struct
    fmt.Printf("All Host info: ", infoStat)
    
}