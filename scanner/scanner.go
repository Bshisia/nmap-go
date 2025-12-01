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
	fmt.Printf("Aggressive scanning %s...\n", host)
	
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		
		if err == nil {
			service := DetectService(conn, port)
			banner := grabBanner(conn)
			fmt.Printf("Port %d: OPEN %s %s\n", port, service, banner)
			conn.Close()
		}
	}
}

func grabBanner(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err == nil && n > 0 {
		banner := strings.TrimSpace(string(buffer[:n]))
		if len(banner) > 50 {
			banner = banner[:50] + "..."
		}
		return fmt.Sprintf("[%s]", banner)
	}
	return ""
}