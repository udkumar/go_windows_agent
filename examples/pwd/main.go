package main

import (
	"fmt"
	"os"
)

func main() {
	// get the present working directory
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Current working directory:", cwd)
	}
}
