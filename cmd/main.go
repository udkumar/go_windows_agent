package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Expand-My-Business/go_windows_agent/netstat"
	"github.com/Expand-My-Business/go_windows_agent/nmap"
	"github.com/Expand-My-Business/go_windows_agent/utils"
	"github.com/Expand-My-Business/go_windows_agent/windowsagent"
	"github.com/sirupsen/logrus"
)

type Message struct {
	data []byte
	url  string
}

func routineANmap(url string, output chan<- Message) {
	nmapXbyte, err := nmap.GetNmapDetails("127.0.0.1", "1-1000")
	if err != nil {
		logrus.Errorf("cannot get nmap details, error: %+v", err)
	}

	for {
		output <- Message{
			data: nmapXbyte,
			url:  url,
		}
		time.Sleep(30 * time.Second)
	}
}

func routineBWindows(url string, output chan<- Message) {
	winXbytes, err := windowsagent.GetWindowsStats()
	if err != nil {
		logrus.Errorf("cannot get windows stats, error: %+v", nil)
	}

	for {
		output <- Message{
			data: winXbytes,
			url:  url,
		}
		time.Sleep(10 * time.Second)
	}
}

func routineCNetStat(url string, output chan<- Message) {
	netXbyte, err := netstat.GetNetStats()
	if err != nil {
		logrus.Errorf("cannot get netstat details, error: %+v", nil)
	}

	for {
		output <- Message{
			data: netXbyte,
			url:  url,
		}
		time.Sleep(10 * time.Second)
	}
}

func Run() {
	currentWD, err := utils.GetWorkingDir()
	if err != nil {
		logrus.Errorf("cannot get the current working dir, error: %+v", err)
	}

	logfilePath := filepath.Join(currentWD, "agent.log")
	// Open a file for writing store the log file in current working directory
	file, err := os.OpenFile(logfilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Errorf("cannot open the logfile, error: %+v", err)
	}
	defer file.Close()

	// Set the log output to the file
	logrus.SetOutput(file)

	fmt.Println("Starting Go routines...")

	// Create channels for communicating with the goroutines
	output := make(chan Message)
	done := make(chan bool)

	// Start the goroutines
	go routineANmap("http://13.235.66.99/agent_ports_data", output)
	go routineBWindows("http://13.235.66.99/add_agent_logs", output)
	go routineCNetStat("http://13.235.66.99/agent_process_data", output)

	// Print the messages from the goroutines as they arrive
	go func() {
		for {
			select {
			case message := <-output:
				fmt.Println("Sending json to the adress :", message.url)
				go sendStringToAPI(message.url, string(message.data))
			case <-done:
				return
			}
		}
	}()

	// Wait for user input to stop the routines
	fmt.Println("Press ENTER to stop the routines.")
	fmt.Scanln()

	// Signal the goroutines to stop
	done <- true

	fmt.Println("Go routines stopped.")
}
func main() {
	Run()
}

func sendStringToAPI(url string, data string) error {
	logrus.Infof("Sending data to API: %s", url)
	requestBody := bytes.NewBuffer([]byte(data))

	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		logrus.Errorf("cannot make a request wrapper, error: %+v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("cannot send a request, error: %+v", err)
		return err
	}
	defer resp.Body.Close()
	logrus.Infof("Ending execution for API: %s", url)
	return nil
}
