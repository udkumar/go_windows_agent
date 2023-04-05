package windowsagent

import (
	"github.com/shirou/gopsutil/cpu"
)

func CpuStats() ([]cpu.InfoStat, error) {
	info, err := cpu.Info()

	//almost every return value is a struct
	// fmt.Printf("All info: ", info)
	return info, err
}
