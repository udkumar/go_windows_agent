package windowslogs

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

func GetSystemLogs() ([]byte, error) {
	secLogs, err := getSecurityLogs()
	if err != nil {
		logrus.Errorf("cannot get security log, error: %v", err)
	}

	appLogs, err := getApplicationLogs()
	if err != nil {
		logrus.Errorf("cannot get application log, error: %v", err)
	}

	logs := Logs{
		SecurityLogs:    secLogs,
		ApplicationLogs: appLogs,
	}
	bx, err := json.MarshalIndent(logs, "", "\t")
	// ioutil.WriteFile("alllogs.json", bx, 0777)
	return bx, err
}

func getSecurityLogs() ([]Log, error) {
	// Calculate the start and end time for the last 2 hours
	startTime := time.Now().Add(-2 * time.Hour).Format("2006-01-02T15:04:05")
	endTime := time.Now().Format("2006-01-02T15:04:05")

	// PowerShell command to retrieve security logs within the time range and convert to JSON
	psCmd := fmt.Sprintf(`Get-WinEvent -FilterHashtable @{
		Logname = 'Security';
		StartTime = '%s';
		EndTime = '%s'
	} | ConvertTo-Json`, startTime, endTime)

	// Run the PowerShell command and capture output and error
	cmd := exec.Command("powershell.exe", "-Command", psCmd)
	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			logrus.Errorf("Security log command failed with error: %v", string(exitError.Stderr))
		} else {
			logrus.Errorf("Failed to execute security log command, error: %v", err)
		}
		return nil, err
	}

	var logs []Log
	err = json.Unmarshal(output, &logs)
	if err != nil {
		logrus.Errorf("cannot unmarshal the logs, error: %v", err)
		return nil, err
	}
	// ioutil.WriteFile("file.json", output, 0777)
	return logs, nil
}

func getApplicationLogs() ([]Log, error) {
	// Calculate the start and end time for the last 2 hours
	startTime := time.Now().Add(-2 * time.Hour).Format("2006-01-02T15:04:05")
	endTime := time.Now().Format("2006-01-02T15:04:05")

	// PowerShell command to retrieve application logs within the time range and convert to JSON
	psCmd := fmt.Sprintf(`Get-WinEvent -FilterHashtable @{
		Logname = 'Application';
		StartTime = '%s';
		EndTime = '%s'
	} | ConvertTo-Json`, startTime, endTime)

	// Run the PowerShell command and capture output and error
	cmd := exec.Command("powershell.exe", "-Command", psCmd)
	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			logrus.Errorf("Application log command failed with error: %v", string(exitError.Stderr))
		} else {
			logrus.Errorf("Failed to execute application log command, error: %v", err)
		}
		return nil, err
	}

	var logs []Log
	err = json.Unmarshal(output, &logs)
	if err != nil {
		logrus.Errorf("cannot unmarshal the logs, error: %v", err)
		return nil, err
	}
	// ioutil.WriteFile("file.json", output, 0777)
	return logs, nil
}
