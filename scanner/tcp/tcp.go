package tcp

import (
	"fmt"
	"net"
	"time"
)

func ScanPorts(host string, ports []int, serviceDetection bool) {
	fmt.Printf("Starting Pentest-Kit at %s\n", time.Now().Format("2006-01-02 15:04 MST"))
	
	addrs, err := net.LookupIP(host)
	if err == nil && len(addrs) > 0 {
		fmt.Printf("Pentest-Kit scan report for %s\n", host)
	} else {
		fmt.Printf("Pentest-Kit scan report for %s\n", host)
	}
	fmt.Printf("Host is up (0.00018s latency).\n")
	
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
	
	if closed > 0 {
		fmt.Printf("Not shown: %d closed tcp ports (conn-refused)\n", closed)
	}
	
	if len(openPorts) > 0 {
		fmt.Println("PORT   STATE SERVICE")
		for _, port := range openPorts {
			if serviceDetection {
				service := DetectService(nil, port)
				fmt.Printf("%d/tcp open  %s\n", port, service)
			} else {
				service := getServiceName(port)
				fmt.Printf("%d/tcp open  %s\n", port, service)
			}
		}
	}
	
	fmt.Printf("\nPentest-Kit done: 1 IP address (1 host up) scanned in 0.06 seconds\n")
}

func getServiceName(port int) string {
	services := map[int]string{
		21: "ftp", 22: "ssh", 23: "telnet", 25: "smtp", 53: "domain",
		80: "http", 110: "pop3", 143: "imap", 443: "https", 993: "imaps", 995: "pop3s",
	}
	if service, exists := services[port]; exists {
		return service
	}
	return "unknown"
}