package windowsagent

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

func DiskStats() (*disk.UsageStat, error) {
	usase, err := disk.Usage("/")

	//almost every return value is a struct
	fmt.Printf("All Usase: ", usase)
	return usase, err
}
