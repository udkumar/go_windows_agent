package netstat

import (
	"encoding/json"

	"github.com/drael/GOnetstat"
	"github.com/sirupsen/logrus"
)

func Netstat() (string, string, error) {
	stats := GOnetstat.Tcp()
	udp_data := GOnetstat.Udp()
	tcp, err := json.MarshalIndent(stats, "", "\t")
	if err != nil {
		logrus.Errorf("cannot get tcp details")
	}
	udp, err := json.MarshalIndent(udp_data, "", "\t")
	if err != nil {
		logrus.Errorf("cannot get tcp details")
	}

	return string(tcp), string(udp), nil
}
