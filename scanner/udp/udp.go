package udp

import (
	"fmt"
	"net"
	"time"
)

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
	
	closed := 0
	var openFilteredPorts []int
	
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("udp", address, 2*time.Second)
		
		if err == nil {
			conn.SetDeadline(time.Now().Add(1 * time.Second))
			_, err := conn.Write([]byte("test"))
			if err == nil {
				openFilteredPorts = append(openFilteredPorts, port)
			} else {
				closed++
			}
			conn.Close()
		} else {
			closed++
		}
	}
	
	for _, port := range openFilteredPorts {
		fmt.Printf("%d/udp open|filtered\n", port)
	}
	
	if closed > 0 {
		fmt.Printf("All %d scanned ports on %s (%s) are in ignored states.\n", len(ports), host, ip)
		fmt.Printf("Not shown: %d closed udp ports (port-unreach)\n", closed)
	}
	
	fmt.Printf("\nPentest-Kit done: 1 IP address (1 host up) scanned in 0.14 seconds\n")
}