package utils

import (
	"strconv"
	"strings"
)

func ParsePortRange(portRange string) []int {
	if portRange == "-" {
		// Scan all ports 1-65535
		var ports []int
		for i := 1; i <= 65535; i++ {
			ports = append(ports, i)
		}
		return ports
	}
	
	var ports []int
	
	if strings.Contains(portRange, "-") {
		parts := strings.Split(portRange, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		
		for i := start; i <= end; i++ {
			ports = append(ports, i)
		}
	} else {
		port, _ := strconv.Atoi(portRange)
		ports = append(ports, port)
	}
	
	return ports
}