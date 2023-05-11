package windowsagent

import (
	"net"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		logrus.Errorf("dial timeout, error: %+v", err)
		return false
	}
	defer conn.Close()
	return true
}
