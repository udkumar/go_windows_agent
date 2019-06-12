package windowsagent

import (
    "fmt"

    "github.com/shirou/gopsutil/host"
)

func HostStats() {
    infoStat, _ := host.Info()

    //almost every return value is a struct
    fmt.Printf("All Host info: ", infoStat)
    return infoStat, err
    
}