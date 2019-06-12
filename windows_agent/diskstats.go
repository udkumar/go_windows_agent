package windowsagent

import (
    "fmt"

    "github.com/shirou/gopsutil/disk"
)

func DiskStats() {
    usase, _ := disk.Usage("/")

    //almost every return value is a struct
    fmt.Printf("All Usase: ", usase)
    return usase, err
}