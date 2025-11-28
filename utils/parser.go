package utils

import (
	"strconv"
	"strings"
)

func ParsePortRange(portRange string) []int {
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