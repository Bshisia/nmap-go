package udp

import (
	"fmt"
	"net"
	"time"
)

func ScanPorts(host string, ports []int) {
	fmt.Printf("UDP scanning %s...\n", host)
	
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("udp", address, 2*time.Second)
		
		if err == nil {
			conn.SetDeadline(time.Now().Add(1 * time.Second))
			_, err := conn.Write([]byte("test"))
			if err == nil {
				fmt.Printf("Port %d/udp: OPEN|FILTERED\n", port)
			}
			conn.Close()
		}
	}
}