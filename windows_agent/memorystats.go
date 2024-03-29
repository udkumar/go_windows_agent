package windowsagent

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

func MemoryStats() (*mem.VirtualMemoryStat, error) {
	v, err := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	fmt.Println(v.String())
	return v, err
}
