package windowslogs

import (
	"encoding/json"
	"fmt"

	"github.com/Expand-My-Business/go_windows_agent/utils"
	"github.com/sirupsen/logrus"
	"github.com/yusufpapurcu/wmi"
)

type Win32_NTEventLogFile struct {
	LogfileName string
}

type LogData struct {
	Data   []Win32_NTLogEvent `json:"logs"`
	HostIP string             `json:"hostIP"`
}

type Win32_NTLogEvent struct {
	Logfile       string `json:"log_file"`
	EventCode     uint16 `json:"event_code"`
	SourceName    string `json:"source_name"`
	Message       string `json:"message"`
	TimeGenerated string `json:"time_generated"`
	EventType     uint16 `json:"event_type"`
}

func GetWindowsEventLog() ([]Win32_NTLogEvent, error) {
	var events []Win32_NTLogEvent

	query := wmi.CreateQuery(&events, "WHERE Logfile='System'")
	query += " AND EventType = 1" // filter by Event Type (Information)
	query += " OR EventType = 2"  // filter by Event Type (Information)
	query += " OR EventType = 3"  // filter by Event Type (Information)
	query += " OR EventType = 5"  // filter by Event Type (Information)

	if err := wmi.Query(query, &events); err != nil {
		logrus.Error("cannot run the log query, error: %+v", err)
		return nil, err
	}

	return events, nil
	// jsonString, err := json.MarshalIndent(events, "", "\t")
	// if err != nil {
	// 	logrus.Error("cannot marshal to even log, error: %+v", err)
	// 	return nil, err
	// }

	// ioutil.WriteFile("eventlog.json", jsonString, 0777)

	// return string(jsonString), nil
}

func GetWindowsApplicationLog() ([]Win32_NTLogEvent, error) {
	var events []Win32_NTLogEvent

	query := wmi.CreateQuery(&events, "WHERE Logfile='Application'")
	if err := wmi.Query(query, &events); err != nil {
		logrus.Error("cannot run the log query, error: %+v", err)
		return nil, err
	}
	return events, nil
	// jsonString, err := json.MarshalIndent(events, "", "\t")
	// if err != nil {
	// 	logrus.Error("cannot marshal to even log, error: %+v", err)
	// 	return "", err
	// }

	// ioutil.WriteFile("application.json", jsonString, 0777)

	// return string(jsonString), nil
}

// GetTypesOfLogs() function retrieves a list of all available logs on the system and prints their names to the console.
// Application
// HardwareEvents
// Internet Explorer
// Key Management Service
// OneApp_IGCC
// Parameters
// Security
// State
// System
// Windows PowerShell
func GetTypesOfLogs() {
	var logs []Win32_NTEventLogFile

	query := wmi.CreateQuery(&logs, "")
	err := wmi.Query(query, &logs)
	if err != nil {
		panic(err)
	}

	for _, log := range logs {
		fmt.Println(log.LogfileName)
	}
}

func GetLogData() ([]byte, error) {
	var logData LogData
	eventLog, err := GetWindowsEventLog()
	if err != nil {
		logrus.Error("cannot run the event logs, error: %+v", err)
	}
	applog, err := GetWindowsApplicationLog()
	if err != nil {
		logrus.Error("cannot run the Application logs, error: %+v", err)
	}
	logData.Data = append(logData.Data, eventLog...)
	logData.Data = append(logData.Data, applog...)

	addr1, err := utils.GetPrivateIPAddress()
	if err != nil {
		logrus.Errorf("cannot get ip address: %+v", err)
	}

	logData.HostIP = addr1

	jsonString, err := json.MarshalIndent(logData, "", "\t")
	if err != nil {
		logrus.Error("cannot marshal to even log, error: %+v", err)
		return nil, err
	}
	// ioutil.WriteFile("sys_logs.json", jsonString, 0777)
	return jsonString, nil
}
