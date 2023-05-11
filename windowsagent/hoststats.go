package windowsagent

import (
	"net"

	"github.com/shirou/gopsutil/host"
	"github.com/sirupsen/logrus"
)

func HostStats() (*host.InfoStat, error) {
	infoStat, err := host.Info()
	if err != nil {
		logrus.Errorf("cannot  get hostInfo, error: %+v", err)
	}
	//almost every return value is a struct
	// logrus.Infof("All Host info: ", infoStat)
	return infoStat, err
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logrus.Errorf("cannot dial to get outbound IP, error: %+v", err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
