package fin

import (
	"fmt"
	"net"
	"time"
)

func ScanPorts(host string, ports []int) {
	fmt.Printf("FIN scanning %s...\n", host)
	
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		
		if err != nil {
			fmt.Printf("Port %d: OPEN|FILTERED\n", port)
		} else {
			conn.Close()
		}
	}
}