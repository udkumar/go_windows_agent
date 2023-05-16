package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func GetPrivateIPAddress() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

// GetWorkingDir get the present working directory
func GetWorkingDir() (string, error) {
	return os.Getwd()
}

func GetPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

func GetMacAddresses() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		logrus.Errorf("cannot check N/W interface, error: %v", err)
		return "", err
	}

	for _, iface := range interfaces {
		// Skip loopback and non-up interfaces
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			logrus.Errorf("cannot get unicast interface, error: %v", err)
			return "", err
		}

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				hwAddr := iface.HardwareAddr
				logrus.Infof("MAC address: %s\n", hwAddr.String())
				return hwAddr.String(), nil
			}
		}
	}

	fmt.Println("MAC address not found")
	return "", errors.New("MAC address not found")
}
