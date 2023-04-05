package netstat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/drael/GOnetstat"
)

type NetStats struct {
	UDPStats []GOnetstat.Process
	TCPStats []GOnetstat.Process
}

func Netstat() (NetStats, error) {
	tcpData := GOnetstat.Tcp()
	udpData := GOnetstat.Udp()

	// tcp, err := json.MarshalIndent(tcpData, "", "\t")
	// if err != nil {
	// 	logrus.Errorf("cannot get tcp details")
	// }

	// udp, err := json.MarshalIndent(udpData, "", "\t")
	// if err != nil {
	// 	logrus.Errorf("cannot get tcp details")
	// }
	netstat := NetStats{
		UDPStats: udpData,
		TCPStats: tcpData,
	}

	bxStats, err := json.MarshalIndent(netstat, "", "    ")
	if err != nil {
		fmt.Errorf("cannot marshal to byteslice", err)
	}
	ioutil.WriteFile("netstat.json", bxStats, 0777)
	return netstat, nil
}
