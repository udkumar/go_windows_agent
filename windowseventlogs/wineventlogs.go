package windowseventlogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/yusufpapurcu/wmi"
)

type Win32_NTEventLogFile struct {
	LogfileName string
}

type Win32_NTLogEvent struct {
	Logfile       string
	EventCode     uint16
	SourceName    string
	Message       string
	TimeGenerated string
	EventType     uint16
}

func GetWindowsEventLog() (string, error) {
	var events []Win32_NTLogEvent

	query := wmi.CreateQuery(&events, "WHERE Logfile='System'")
	query += " AND EventType = 1" // filter by Event Type (Information)
	query += " OR EventType = 2"  // filter by Event Type (Information)
	query += " OR EventType = 3"  // filter by Event Type (Information)
	query += " OR EventType = 5"  // filter by Event Type (Information)

	if err := wmi.Query(query, &events); err != nil {
		logrus.Error("cannot run the log query, error: %+v", err)
		return "", err
	}

	jsonString, err := json.MarshalIndent(events, "", "\t")
	if err != nil {
		logrus.Error("cannot marshal to even log, error: %+v", err)
		return "", err
	}

	ioutil.WriteFile("eventlog.json", jsonString, 0777)

	return string(jsonString), nil
}

func GetWindowsApplicationLog() (string, error) {
	var events []Win32_NTLogEvent

	query := wmi.CreateQuery(&events, "WHERE Logfile='Application'")
	err := wmi.Query(query, &events)

	if err := wmi.Query(query, &events); err != nil {
		logrus.Error("cannot run the log query, error: %+v", err)
		return "", err
	}

	jsonString, err := json.MarshalIndent(events, "", "\t")
	if err != nil {
		logrus.Error("cannot marshal to even log, error: %+v", err)
		return "", err
	}

	ioutil.WriteFile("application.json", jsonString, 0777)

	return string(jsonString), nil
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
