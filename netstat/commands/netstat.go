package commands

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type NetstatCmd struct {
	NetstatA []NetstatA `json:"netstat_a"`
	NetstatB []NetstatB `json:"netstat_b"`
}

type NetstatA struct {
	Proto          string `json:"proto"`
	LocalAddress   string `json:"local_address"`
	ForeignAddress string `json:"foreign_address"`
	State          string `json:"state"`
}
type NetstatB struct {
	Proto          string `json:"proto"`
	LocalAddress   string `json:"local_address"`
	ForeignAddress string `json:"foreign_address"`
	State          string `json:"state"`
	ProcessName    string `json:"process_name"`
}
type Process struct {
	ImageName   string
	PID         int
	SessionName string
	SessionNum  int
	MemUsage    string
	Status      string
	UserName    string
	CPUTime     string
	WindowTitle string
}

func GetNetstatA() ([]NetstatA, error) {
	// Run the netstat command and capture its output
	out, err := exec.Command("netstat", "-a").Output()
	if err != nil {
		logrus.Errorf("cannot run netstat -a command, error: %v", err)
		return nil, err
	}

	// Split the output into lines and discard the first two lines
	lines := strings.Split(string(out), "\n")[2:]

	// Parse each line into a NetstatA struct and add it to the result array
	var result []NetstatA
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			result = append(result, NetstatA{
				Proto:          fields[0],
				LocalAddress:   fields[1],
				ForeignAddress: fields[2],
				State:          fields[3],
			})
		}
	}
	return result, nil
	// // Print the result as JSON
	// netstatA, err := json.MarshalIndent(result, "", "\t")
	// if err != nil {
	// 	logrus.Errorf("cannot unmarshal netstat -a values, error: %v", err)
	// 	return nil, err
	// }

	// // ioutil.WriteFile("netstata.json", netstatA, 0777)
	// // fmt.Println(string(netstatA))
	// return netstatA, nil
}

func GetNetstatB() ([]NetstatB, error) {
	// Run the netstat command and capture its output
	out, err := exec.Command("netstat", "-b").Output()
	if err != nil {
		logrus.Errorf("cannot run netstat -b command, error: %v", err)
		return nil, err
	}

	// Split the output into lines and discard the first three lines
	lines := strings.Split(string(out), "\n")[3:]

	// Parse each line into a NetstatB struct and add it to the result array
	var result []NetstatB
	for i := 0; i < len(lines); i += 2 {
		header := strings.Fields(lines[i])
		processName := header[len(header)-1]

		data := strings.Fields(lines[i+1])
		if len(data) >= 4 {
			result = append(result, NetstatB{
				Proto:          data[0],
				LocalAddress:   data[1],
				ForeignAddress: data[2],
				State:          data[3],
				ProcessName:    processName,
			})
		}
	}
	return result, nil

	// // Print the result as JSON
	// netstatB, err := json.MarshalIndent(result, "", "\t")
	// if err != nil {
	// 	logrus.Errorf("cannot unmarshal netstat -b values, error: %v", err)
	// 	return nil, err
	// }
	// // ioutil.WriteFile("netstatb.json", jsonResult, 777)
	// // fmt.Println(string(jsonResult))

	// return netstatB, nil
}

func GetNetStats() {
	nsa, _ := GetNetstatA()
	nsb, _ := GetNetstatB()

	nsCmd := NetstatCmd{
		NetstatA: nsa,
		NetstatB: nsb,
	}

	nsCmdByte, err := json.MarshalIndent(nsCmd, "", "\t")
	if err != nil {
		logrus.Errorf("cannot unmarshal netstat -b values, error: %v", err)
	}
	ioutil.WriteFile("netstatcmd.json", nsCmdByte, 777)
}

func GetAllInternalProcess() ([]Process, error) {
	cmd := exec.Command("tasklist", "/v", "/fo", "csv")
	output, err := cmd.Output()
	if err != nil {
		logrus.Errorf("cannot run task command, error: %v", err)
		return nil, err
	}
	reader := csv.NewReader(strings.NewReader(string(output)))
	records, err := reader.ReadAll()
	if err != nil {
		logrus.Errorf("cannot read tasklist output, error: %v", err)
		return nil, err
	}
	processes := make([]Process, len(records)-1)
	for i, record := range records[1:] {
		pid, _ := strconv.Atoi(record[1])
		sessionNum, _ := strconv.Atoi(record[3])
		processes[i] = Process{
			ImageName:   record[0],
			PID:         pid,
			SessionName: record[2],
			SessionNum:  sessionNum,
			MemUsage:    record[4],
			Status:      record[5],
			UserName:    record[6],
			CPUTime:     record[7],
			WindowTitle: record[8],
		}
	}

	// processesByte, err := json.MarshalIndent(processes, "", "\t")
	// if err != nil {
	// 	logrus.Errorf("cannot unmarshal unmarshal processes, error: %v", err)
	// 	return nil, err
	// }
	// ioutil.WriteFile("allinternalprocess.json", processesByte, 0777)

	return processes, nil
}
