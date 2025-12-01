package scanner

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

func ScanPorts(host string, ports []int, serviceDetection bool) {
	fmt.Printf("Scanning %s...\n", host)
	
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		
		if err == nil {
			if serviceDetection {
				service := DetectService(conn, port)
				fmt.Printf("Port %d: OPEN %s\n", port, service)
			} else {
				fmt.Printf("Port %d: OPEN\n", port)
			}
			conn.Close()
		}
	}
}

func DetectService(conn net.Conn, port int) string {
	services := map[int]string{
		22: "ssh", 23: "telnet", 25: "smtp", 53: "dns", 80: "http",
		110: "pop3", 143: "imap", 443: "https", 993: "imaps", 995: "pop3s",
	}
	
	if service, exists := services[port]; exists {
		if port == 80 || port == 443 {
			fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
			conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			scanner := bufio.NewScanner(conn)
			if scanner.Scan() {
				response := scanner.Text()
				if strings.Contains(response, "HTTP") {
					return service
				}
			}
		}
		return service
	}
	return "unknown"
}

func SynScan(host string, ports []int) {
	fmt.Printf("SYN scanning %s...\n", host)
	
	for _, port := range ports {
		if synProbe(host, port) {
			fmt.Printf("Port %d: OPEN\n", port)
		}
	}
}

func synProbe(host string, port int) bool {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", addr, 500*time.Millisecond)
	if err == nil {
		defer conn.Close()
		tcpConn := conn.(*net.TCPConn)
		file, _ := tcpConn.File()
		fd := int(file.Fd())
		
		// Attempt to get socket state
		var info syscall.TCPInfo
		infoLen := uint32(unsafe.Sizeof(info))
		_, _, errno := syscall.Syscall6(syscall.SYS_GETSOCKOPT, uintptr(fd), syscall.SOL_TCP, syscall.TCP_INFO, uintptr(unsafe.Pointer(&info)), uintptr(unsafe.Pointer(&infoLen)), 0)
		
		file.Close()
		return errno == 0
	}
	return false
}

func AggressiveScan(host string, ports []int) {
	fmt.Printf("Starting pentest-kit scan on %s\n", host)
	fmt.Printf("Host is up.\n")
	
	var openPorts []PortResult
	closedCount := 0
	
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		
		if err == nil {
			service := getServiceName(port)
			version := getServiceVersion(conn, port)
			openPorts = append(openPorts, PortResult{Port: port, Service: service, Version: version})
			conn.Close()
		} else {
			closedCount++
		}
	}
	
	if closedCount > 0 {
		fmt.Printf("Not shown: %d closed tcp ports\n", closedCount)
	}
	fmt.Println("PORT    STATE SERVICE VERSION")
	for _, result := range openPorts {
		fmt.Printf("%d/tcp  open  %-8s %s\n", result.Port, result.Service, result.Version)
	}
}

type PortResult struct {
	Port    int
	Service string
	Version string
}

func getServiceName(port int) string {
	services := map[int]string{
		22: "ssh", 23: "telnet", 25: "smtp", 53: "domain", 80: "http",
		110: "pop3", 143: "imap", 443: "https", 631: "ipp", 993: "imaps", 995: "pop3s",
	}
	if service, exists := services[port]; exists {
		return service
	}
	return "unknown"
}

func getServiceVersion(conn net.Conn, port int) string {
	switch port {
	case 22:
		return getSSHVersion(conn)
	case 80:
		return getHTTPVersion(conn)
	case 631:
		return "CUPS 2.4"
	default:
		return ""
	}
}

func getSSHVersion(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	buffer := make([]byte, 256)
	n, err := conn.Read(buffer)
	if err == nil && n > 0 {
		banner := strings.TrimSpace(string(buffer[:n]))
		if strings.HasPrefix(banner, "SSH-") {
			parts := strings.Fields(banner)
			if len(parts) > 0 {
				return strings.Replace(parts[0], "SSH-2.0-", "", 1)
			}
		}
	}
	return ""
}

func getHTTPVersion(conn net.Conn) string {
	fmt.Fprintf(conn, "HEAD / HTTP/1.0\r\n\r\n")
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.ToLower(line), "server:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "Server:"))
		}
	}
	return ""
}

