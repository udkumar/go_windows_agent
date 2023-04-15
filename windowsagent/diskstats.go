package windowsagent

import (
	"github.com/shirou/gopsutil/disk"
)

func DiskStats() (*disk.UsageStat, error) {
	usase, err := disk.Usage("/")

	//almost every return value is a struct
	// logrus.Errorf("All Usase: %v", usase)
	return usase, err
}
