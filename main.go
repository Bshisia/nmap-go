package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <host> <port-range>")
		fmt.Println("Example: go run main.go 192.168.1.1 80-443")
		os.Exit(1)
	}

}
