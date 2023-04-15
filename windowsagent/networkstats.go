package windowsagent

import (
	"github.com/shirou/gopsutil/net"
	"github.com/sirupsen/logrus"
)

func NetStats() ([]net.IOCountersStat, error) {
	infoIOCounter, err := net.IOCounters(true)
	if err != nil {
		logrus.Errorf("cannot get IOCounters, error: %+v", err)
	}
	return infoIOCounter, err
}
