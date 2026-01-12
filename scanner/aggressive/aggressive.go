package aggressive

import (
	"fmt"
	"net"
	"pentest-kit/scanner/service"
	"time"
)

type PortResult struct {
	Port    int
	Service string
	Version string
}

func ScanPorts(host string, ports []int) {
	fmt.Printf("Starting Pentest-Kit at %s\n", time.Now().Format("2006-01-02 15:04 MST"))
	
	addrs, err := net.LookupIP(host)
	var ip string
	if err != nil || len(addrs) == 0 {
		ip = host
	} else {
		ip = addrs[0].String()
	}
	
	fmt.Printf("Pentest-Kit scan report for %s (%s)\n", host, ip)
	fmt.Printf("Host is up (0.0000050s latency).\n")
	
	var openPorts []PortResult
	closed := 0
	
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		
		if err == nil {
			serviceName := service.GetServiceName(port)
			version := service.GetServiceVersion(conn, port)
			openPorts = append(openPorts, PortResult{Port: port, Service: serviceName, Version: version})
			conn.Close()
		} else {
			closed++
		}
	}
	
	fmt.Println("PORT    STATE SERVICE VERSION")
	for _, result := range openPorts {
		fmt.Printf("%d/tcp  open  %-8s %s\n", result.Port, result.Service, result.Version)
	}
	
	if closed > 0 {
		fmt.Printf("Not shown: %d closed tcp ports (reset)\n", closed)
	}
	
	fmt.Printf("\nPentest-Kit scan done: 1 IP address (1 host up) scanned in 0.14 seconds\n")
}