package main

import (
	"fmt"
	"time"
)

func routineA(output chan<- string) {
	for {
		output <- "Routine A is running."
		time.Sleep(10 * time.Second)
	}
}

func routineB(output chan<- string) {
	for {
		output <- "Routine B is running."
		time.Sleep(1 * time.Second)
	}
}

func routineC(output chan<- string) {
	for {
		output <- "Routine C is running."
		time.Sleep(5 * time.Second)
	}
}

func main() {
	fmt.Println("Starting Go routines...")

	// Create channels for communicating with the goroutines
	output := make(chan string)
	done := make(chan bool)

	// Start the goroutines
	go routineA(output)
	go routineB(output)
	go routineC(output)

	// Print the messages from the goroutines as they arrive
	go func() {
		for {
			select {
			case message := <-output:
				fmt.Println(message)
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
