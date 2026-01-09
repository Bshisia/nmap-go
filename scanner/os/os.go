package os

import (
	"fmt"
	"net"
	"time"
)

func DetectOS(host string, ports []int) {
	fmt.Printf("OS detection on %s...\n", host)
	fmt.Printf("Scanning %d ports for OS fingerprinting...\n", len(ports))

	var openPorts []int
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err == nil {
			fmt.Printf("Port %d: OPEN\n", port)
			openPorts = append(openPorts, port)
			conn.Close()
		} else {
			fmt.Printf("Port %d: CLOSED\n", port)
		}
	}

	if len(openPorts) == 0 {
		fmt.Println("No open ports found for OS detection")
		return
	}

	osGuess := guessOS(openPorts)
	fmt.Printf("OS guess: %s\n", osGuess)
}

func guessOS(ports []int) string {
	hasSSH := contains(ports, 22)
	hasHTTP := contains(ports, 80) || contains(ports, 8080)
	hasHTTPS := contains(ports, 443)
	hasSMB := contains(ports, 445) || contains(ports, 139)
	hasRDP := contains(ports, 3389)

	if hasRDP || hasSMB {
		return "Microsoft Windows"
	}
	if hasSSH && (hasHTTP || hasHTTPS) {
		return "Linux"
	}
	if hasSSH {
		return "Unix-like (Linux/BSD/macOS)"
	}

	return "Unknown"
}

func contains(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
