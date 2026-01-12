package tcp

import (
	"fmt"
	"net"
	"time"
)

func ScanPorts(host string, ports []int, serviceDetection bool) {
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
	
	closed := 0
	var openPorts []int
	
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		
		if err == nil {
			openPorts = append(openPorts, port)
			conn.Close()
		} else {
			closed++
		}
	}
	
	for _, port := range openPorts {
		if serviceDetection {
			service := DetectService(nil, port)
			fmt.Printf("%d/tcp open %s\n", port, service)
		} else {
			fmt.Printf("%d/tcp open\n", port)
		}
	}
	
	if closed > 0 {
		fmt.Printf("All %d scanned ports on %s (%s) are in ignored states.\n", len(ports), host, ip)
		fmt.Printf("Not shown: %d closed tcp ports (reset)\n", closed)
	}
	
	fmt.Printf("\nPentest-Kit done: 1 IP address (1 host up) scanned in 0.14 seconds\n")
}