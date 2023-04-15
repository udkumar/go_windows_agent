package utils

import (
	"net"
	"os"
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
