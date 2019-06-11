package main

import (
    "fmt"

    "github.com/shirou/gopsutil/disk"
)

func main() {
    usase, _ := disk.Usage("/")

    //almost every return value is a struct
    fmt.Printf("All Usase: ", usase)
    
}