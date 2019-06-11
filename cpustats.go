package cpu

import (
    "fmt"

    "github.com/shirou/gopsutil/cpu"
)

func CpuStats() {
    info, _ := cpu.Info()

    //almost every return value is a struct
    fmt.Printf("All info: ", info)
    
}