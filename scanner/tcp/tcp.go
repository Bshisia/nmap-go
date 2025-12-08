package tcp

import (
	"fmt"
	"net"
	"time"
)

func ScanPorts(host string, ports []int, serviceDetection bool) {
	fmt.Printf("Scanning %s...\n", host)
	
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		
		if err == nil {
			if serviceDetection {
				service := DetectService(conn, port)
				fmt.Printf("Port %d: OPEN %s\n", port, service)
			} else {
				fmt.Printf("Port %d: OPEN\n", port)
			}
			conn.Close()
		}
	}
}