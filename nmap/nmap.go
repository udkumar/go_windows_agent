package nmap

import (
	"encoding/json"

	"github.com/Expand-My-Business/go_windows_agent/utils"
	"github.com/Ullaakut/nmap"
	"github.com/sirupsen/logrus"
)

type NmapDetails struct {
	Nmap   NmapStats `json:"sys_ports"`
	HostIP string    `json:"hostIP"`
}

type NmapStats []struct {
	Distance struct {
		Value int `json:"value"`
	} `json:"distance"`
	EndTime      int `json:"end_time"`
	IPIDSequence struct {
		Class  string `json:"class"`
		Values string `json:"values"`
	} `json:"ip_id_sequence"`
	Os struct {
		PortsUsed      interface{} `json:"ports_used"`
		OsMatches      interface{} `json:"os_matches"`
		OsFingerprints interface{} `json:"os_fingerprints"`
	} `json:"os"`
	StartTime int `json:"start_time"`
	Status    struct {
		State     string `json:"state"`
		Reason    string `json:"reason"`
		ReasonTTL int    `json:"reason_ttl"`
	} `json:"status"`
	TCPSequence struct {
		Index      int    `json:"index"`
		Difficulty string `json:"difficulty"`
		Values     string `json:"values"`
	} `json:"tcp_sequence"`
	TCPTsSequence struct {
		Class  string `json:"class"`
		Values string `json:"values"`
	} `json:"tcp_ts_sequence"`
	Times struct {
		Srtt string `json:"srtt"`
		Rttv string `json:"rttv"`
		To   string `json:"to"`
	} `json:"times"`
	Trace struct {
		Proto string      `json:"proto"`
		Port  int         `json:"port"`
		Hops  interface{} `json:"hops"`
	} `json:"trace"`
	Uptime struct {
		Seconds  int    `json:"seconds"`
		LastBoot string `json:"last_boot"`
	} `json:"uptime"`
	Comment   string `json:"comment"`
	Addresses []struct {
		Addr     string `json:"addr"`
		AddrType string `json:"addr_type"`
		Vendor   string `json:"vendor"`
	} `json:"addresses"`
	ExtraPorts []struct {
		State   string `json:"state"`
		Count   int    `json:"count"`
		Reasons []struct {
			Reason string `json:"reason"`
			Count  int    `json:"count"`
		} `json:"reasons"`
	} `json:"extra_ports"`
	Hostnames []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"hostnames"`
	HostScripts interface{} `json:"host_scripts"`
	Ports       []struct {
		ID       int    `json:"id"`
		Protocol string `json:"protocol"`
		Owner    struct {
			Name string `json:"name"`
		} `json:"owner"`
		Service struct {
			DeviceType    string   `json:"device_type"`
			ExtraInfo     string   `json:"extra_info"`
			HighVersion   string   `json:"high_version"`
			Hostname      string   `json:"hostname"`
			LowVersion    string   `json:"low_version"`
			Method        string   `json:"method"`
			Name          string   `json:"name"`
			OsType        string   `json:"os_type"`
			Product       string   `json:"product"`
			Proto         string   `json:"proto"`
			RPCNum        string   `json:"rpc_num"`
			ServiceFp     string   `json:"service_fp"`
			Tunnel        string   `json:"tunnel"`
			Version       string   `json:"version"`
			Configuration int      `json:"configuration"`
			Cpes          []string `json:"cpes"`
		} `json:"service"`
		State struct {
			State     string `json:"state"`
			Reason    string `json:"reason"`
			ReasonIP  string `json:"reason_ip"`
			ReasonTTL int    `json:"reason_ttl"`
		} `json:"state"`
		Scripts interface{} `json:"scripts"`
	} `json:"ports"`
	Smurfs interface{} `json:"smurfs"`
}

func GetNmapDetails(addr, portRange string) ([]byte, error) {
	// Create a new Nmap scanner instance
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(addr),
		nmap.WithPorts(portRange),
		nmap.WithServiceInfo(),
		nmap.WithVersionAll(),
		nmap.WithVersionTrace(),
	)
	if err != nil {
		logrus.Fatalf("failed to create scanner: %v", err)
		return nil, err
	}

	// Run the scan
	result, _, err := scanner.Run()
	if err != nil {
		logrus.Fatalf("failed to run scan: %v", err)
		return nil, err
	}

	// Print any warnings from the scan
	// logrus.Warnf("Warnings: %+v\n", warnings)

	bx, err := json.MarshalIndent(result.Hosts, "", "\t")
	if err != nil {
		logrus.Errorf("cannot marshal json: %+v", err)
		return nil, err
	}

	ns := &NmapStats{}
	if err = json.Unmarshal(bx, ns); err != nil {
		logrus.Errorf("cannot unmarshal to json: %+v", err)
		return nil, err
	}

	addr1, err := utils.GetPrivateIPAddress()
	if err != nil {
		logrus.Errorf("cannot get ip address: %+v", err)
		return nil, err
	}

	pd := NmapDetails{
		Nmap:   *ns,
		HostIP: addr1,
	}

	bxPd, err := json.MarshalIndent(pd, "", "\t")
	if err != nil {
		logrus.Errorf("cannot marshal json: %+v", err)
		return nil, err
	}

	// if err := ioutil.WriteFile("old_nmap.json", bxPd, 0644); err != nil {
	// 	panic(err)
	// }
	return bxPd, nil
}
