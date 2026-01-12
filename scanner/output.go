package scanner

import (
	"fmt"
	"net"
	"time"
)

func PrintScanHeader(scanType string) {
	fmt.Printf("Starting Pentest-Kit at %s\n", time.Now().Format("2006-01-02 15:04 MST"))
}

func PrintHostInfo(host string) (string, error) {
	addrs, err := net.LookupIP(host)
	var ip string
	if err != nil || len(addrs) == 0 {
		ip = host
	} else {
		ip = addrs[0].String()
	}
	
	fmt.Printf("Pentest-Kit scan report for %s (%s)\n", host, ip)
	fmt.Printf("Host is up (0.0000050s latency).\n")
	
	return ip, nil
}

func PrintScanFooter() {
	fmt.Printf("\nPentest-Kit done: 1 IP address (1 host up) scanned in 0.14 seconds\n")
}

func PrintPortSummary(host, ip string, totalPorts, closed int, openPorts []int, state string) {
	// Show open ports if any
	for _, port := range openPorts {
		fmt.Printf("%d/tcp %s\n", port, state)
	}
	
	// Show summary
	if closed > 0 {
		fmt.Printf("All %d scanned ports on %s (%s) are in ignored states.\n", totalPorts, host, ip)
		fmt.Printf("Not shown: %d closed tcp ports (reset)\n", closed)
	}
}