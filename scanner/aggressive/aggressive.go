package aggressive

import (
	"fmt"
	"net"
	"pentest-kit/scanner"
	"pentest-kit/scanner/service"
	"time"
)

func ScanPorts(host string, ports []int) {
	fmt.Printf("Starting pentest-kit scan on %s\n", host)
	fmt.Printf("Host is up.\n")
	
	var openPorts []scanner.PortResult
	closedCount := 0
	
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		
		if err == nil {
			serviceName := service.GetServiceName(port)
			version := service.GetServiceVersion(conn, port)
			openPorts = append(openPorts, scanner.PortResult{Port: port, Service: serviceName, Version: version})
			conn.Close()
		} else {
			closedCount++
		}
	}
	
	if closedCount > 0 {
		fmt.Printf("Not shown: %d closed tcp ports\n", closedCount)
	}
	fmt.Println("PORT    STATE SERVICE VERSION")
	for _, result := range openPorts {
		fmt.Printf("%d/tcp  open  %-8s %s\n", result.Port, result.Service, result.Version)
	}
}