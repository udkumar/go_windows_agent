package windowsagent

import (
	"fmt"
    "net"
	"github.com/shirou/gopsutil/host"
)

func HostStats() (*host.InfoStat, error) {
	infoStat, err := host.Info()

	//almost every return value is a struct
	fmt.Printf("All Host info: ", infoStat)
	return infoStat, err
}


// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        fmt.Printf("err: ",err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Printf("localAddr info: ", localAddr)
    return localAddr.IP
}