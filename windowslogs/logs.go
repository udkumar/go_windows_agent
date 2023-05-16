package windowslogs

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

func GetSecurityLogs() (string, error) {
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
			logrus.Errorf("Command failed with error: %v", string(exitError.Stderr))
		} else {
			logrus.Errorf("Failed to execute command, error: %v", err)
		}
		return "", err
	}

	ioutil.WriteFile("file.json", output, 0777)
	return string(output), nil
}
