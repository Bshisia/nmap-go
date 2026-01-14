package tools

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func HostMapper(subnet string, output string) {
	results := []string{}
	results = append(results, fmt.Sprintf("Network mapping: %s", subnet))
	results = append(results, "")
	
	hosts := parseSubnet(subnet)
	
	for _, host := range hosts {
		if isHostAlive(host) {
			result := fmt.Sprintf("%s - ALIVE", host)
			results = append(results, result)
			fmt.Println(result)
		}
	}
	
	if output != "" {
		writeToFile(output, results)
	}
}

func parseSubnet(subnet string) []string {
	hosts := []string{}
	parts := strings.Split(subnet, "/")
	
	if len(parts) != 2 {
		return hosts
	}
	
	baseIP := parts[0]
	ipParts := strings.Split(baseIP, ".")
	
	if len(ipParts) != 4 {
		return hosts
	}
	
	for i := 1; i <= 254; i++ {
		host := fmt.Sprintf("%s.%s.%s.%d", ipParts[0], ipParts[1], ipParts[2], i)
		hosts = append(hosts, host)
	}
	
	return hosts
}

func isHostAlive(host string) bool {
	timeout := 500 * time.Millisecond
	
	// Try ICMP-like check via TCP connection to common ports
	ports := []int{80, 443, 22, 21, 23}
	
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err == nil {
			conn.Close()
			return true
		}
	}
	
	return false
}

func parseIP(ip string) ([]int, error) {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return nil, fmt.Errorf("invalid IP")
	}
	
	result := make([]int, 4)
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		result[i] = num
	}
	return result, nil
}
