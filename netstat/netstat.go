package netstat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/drael/GOnetstat"
	"github.com/sirupsen/logrus"
)

type NetStats struct {
	UDPStats []GOnetstat.Process
	TCPStats []GOnetstat.Process
	HostIP   string `json:"hostIP"`
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
	addr1, err := getIPAddress()
	if err != nil {
		logrus.Errorf("cannot get ip address: %+v", err)
		return NetStats{}, err
	}

	netstat := NetStats{
		UDPStats: udpData,
		TCPStats: tcpData,
		HostIP:   addr1,
	}

	bxStats, err := json.MarshalIndent(netstat, "", "    ")
	if err != nil {
		fmt.Errorf("cannot marshal to byteslice", err)
	}
	ioutil.WriteFile("netstat.json", bxStats, 0777)
	return netstat, nil
}

func getIPAddress() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
