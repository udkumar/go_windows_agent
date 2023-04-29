package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Expand-My-Business/go_windows_agent/netstat"
	"github.com/Expand-My-Business/go_windows_agent/nmap"
	"github.com/Expand-My-Business/go_windows_agent/windowsagent"
	"github.com/kardianos/service"
	"github.com/sirupsen/logrus"
)

type Message struct {
	data []byte
	url  string
	err  error
}

func routineANmap(url string, output chan<- Message, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			nmapXbyte, err := nmap.GetNmapDetails("127.0.0.1", "1-65000")
			if err != nil {
				logrus.Errorf("cannot get nmap details, error: %+v", err)
				output <- Message{
					err: err,
					url: url,
				}
			} else {
				output <- Message{
					data: nmapXbyte,
					url:  url,
				}
			}
			time.Sleep(30 * time.Second)
		}
	}
}

func routineBWindows(url string, output chan<- Message, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			winXbytes, err := windowsagent.GetWindowsStats()
			if err != nil {
				logrus.Errorf("cannot get windows stats, error: %+v", err)
				output <- Message{
					err: err,
					url: url,
				}
			} else {
				output <- Message{
					data: winXbytes,
					url:  url,
				}
			}
			time.Sleep(10 * time.Second)
		}
	}
}

func routineCNetStat(url string, output chan<- Message, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			netXbyte, err := netstat.GetNetStats()
			if err != nil {
				logrus.Errorf("cannot get netstat details, error: %+v", err)
				output <- Message{
					err: err,
					url: url,
				}
			} else {
				output <- Message{
					data: netXbyte,
					url:  url,
				}
			}
			time.Sleep(10 * time.Second)
		}
	}
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

// Define a struct that implements the service.Service interface.
type myService struct {
	done chan bool
}

// Implement the required service methods.
func (m *myService) Start(s service.Service) error {
	m.done = make(chan bool)
	go m.run()
	return nil
}

func (m *myService) Stop(s service.Service) error {
	// Signal the goroutines to stop by closing the channels
	close(m.done)
	return nil
}

func (m *myService) run() {
	fmt.Println("Starting Go routines...")

	// Create channels for communicating with the goroutines
	output := make(chan Message)

	// Start the goroutines
	go routineANmap("http://13.235.66.99/agent_ports_data", output, nil)
	go routineBWindows("http://13.235.66.99/add_agent_logs", output, nil)
	go routineCNetStat("http://13.235.66.99/agent_process_data", output, nil)

	// Print the messages from the goroutines as they arrive
	go func() {
		for {
			select {
			case message := <-output:
				fmt.Println("Sending json to the adress: ", message.url)
				go sendStringToAPI(message.url, string(message.data))
			case <-m.done:
				// Stop the goroutines by closing the output channel and waiting for them to finish
				close(output)
				return
			}
		}
	}()

	// // Wait for user input to stop the service
	// fmt.Println("Press ENTER to stop the service.")
	// fmt.Scanln()

	// // Stop the service by calling service.Stop
	// err := service.Control(m, service)
	// if err != nil {
	// 	logrus.Errorf("Error stopping the service: %s\n", err)
	// }
}

func main() {
	// Check if directory exists, if not create it
	folderPath := `C:\Program Files\GoAgent`
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Call your existing service function here.
	file, err := os.OpenFile(folderPath+"\\go_agent.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Errorf("cannot open the logfile, error: %+v", err)
	}
	defer file.Close()

	// Set the log output to the file
	logrus.SetOutput(file)

	// Create a new service object and initialize it.
	svcConfig := &service.Config{
		Name:        "Agent Service",
		DisplayName: "Agent Service",
		Description: "My service description.",
	}

	prg := &myService{}
	svc, err := service.New(prg, svcConfig)
	if err != nil {
		logrus.Errorf("Error: %s\n", err)
		return
	}

	// Start the service. If the service is already running, this call will block until the service stops.
	err = svc.Run()
	if err != nil {
		logrus.Errorf("Error: %s\n", err)
		return
	}
}
