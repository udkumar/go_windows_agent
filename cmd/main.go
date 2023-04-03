package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	Winticker := time.NewTicker(1 * time.Second)
	Nmapticker := time.NewTicker(3 * time.Second)

	go Nmap(c, Nmapticker)
	go Windows(c, Winticker)

	for {
		select {
		case val := <-c:
			fmt.Printf("Received value from channel: %d\n", val)

		}
		fmt.Println()
	}
}

func Nmap(c chan<- int, ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			fmt.Println("Nmap sends value to channel")
			c <- 1
		}
	}
}

func Windows(c chan<- int, ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			fmt.Println("Windows sends value to channel")
			c <- 2
		}
	}
}
