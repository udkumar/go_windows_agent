package windowsagent

import (
	"github.com/shirou/gopsutil/mem"
	"github.com/sirupsen/logrus"
)

func MemoryStats() (*mem.VirtualMemoryStat, error) {
	virtualMem, err := mem.VirtualMemory()
	if err != nil {
		logrus.Errorf("cannot dial to get outbound IP, error: %+v", err)
	}
	return virtualMem, err
}
