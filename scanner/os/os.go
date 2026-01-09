package os

import (
	"fmt"
	"net"
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

type OSFingerprint struct {
	TTL           int
	WindowSize    int
	TCPOptions    []byte
	ResponseFlags int
}

func DetectOS(host string, ports []int) {
	fmt.Printf("Starting pentest-kit OS detection on %s\n", host)
	fmt.Printf("Scanning %d ports for OS fingerprinting...\n\n", len(ports))

	var openPorts []int
	closedCount := 0

	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 200*time.Millisecond)
		if err == nil {
			fmt.Printf("Port %d: OPEN\n", port)
			openPorts = append(openPorts, port)
			conn.Close()
		} else {
			closedCount++
		}
	}

	if closedCount > 0 {
		fmt.Printf("\nNot shown: %d closed ports\n", closedCount)
	}

	if len(openPorts) == 0 {
		fmt.Println("\nNo open ports found for OS detection")
		fmt.Println("OS guess: Unknown (insufficient data)")
		return
	}

	// Perform TCP fingerprinting
	fingerprint := performTCPFingerprinting(host, openPorts[0])
	osGuess := analyzeFingerprint(fingerprint, openPorts)
	
	fmt.Printf("\nTCP fingerprinting results:\n")
	fmt.Printf("TTL: %d\n", fingerprint.TTL)
	fmt.Printf("Window Size: %d\n", fingerprint.WindowSize)
	fmt.Printf("\nOS guess: %s\n", osGuess)
}

func performTCPFingerprinting(host string, port int) OSFingerprint {
	fingerprint := OSFingerprint{TTL: 64, WindowSize: 65535}
	
	// Try to get more detailed TCP info using raw sockets
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_TCP)
	if err == nil {
		defer syscall.Close(fd)
		
		// Resolve host
		addr, err := net.ResolveIPAddr("ip4", host)
		if err == nil {
			// Send SYN packet and analyze response
			packet := createSynPacket(uint16(port))
			sockaddr := &syscall.SockaddrInet4{Port: port}
			copy(sockaddr.Addr[:], addr.IP.To4())
			
			syscall.Sendto(fd, packet, 0, sockaddr)
			
			// Read response
			buf := make([]byte, 1024)
			syscall.SetNonblock(fd, true)
			time.Sleep(100 * time.Millisecond)
			
			n, _, err := syscall.Recvfrom(fd, buf, 0)
			if err == nil && n > 20 {
				// Parse IP header for TTL
				if len(buf) > 8 {
					fingerprint.TTL = int(buf[8])
				}
				// Parse TCP header for window size
				if len(buf) > 34 {
					fingerprint.WindowSize = int(buf[34])<<8 | int(buf[35])
				}
			}
		}
	}
	
	return fingerprint
}

func createSynPacket(dstPort uint16) []byte {
	header := struct {
		Src    uint16
		Dst    uint16
		Seq    uint32
		Ack    uint32
		Flags  uint16
		Window uint16
		Csum   uint16
		Urg    uint16
	}{
		Src:    12345,
		Dst:    dstPort,
		Seq:    0,
		Ack:    0,
		Flags:  0x5002, // SYN flag + header length
		Window: 65535,
		Csum:   0,
		Urg:    0,
	}
	
	return (*(*[20]byte)(unsafe.Pointer(&header)))[:]
}

func analyzeFingerprint(fp OSFingerprint, openPorts []int) string {
	// Analyze TTL values (common OS signatures)
	switch {
	case fp.TTL <= 64:
		// Linux/Unix systems typically use TTL 64
		if hasLinuxPorts(openPorts) {
			return "Linux"
		}
		return "Unix-like (Linux/BSD/macOS)"
	case fp.TTL <= 128:
		// Windows systems typically use TTL 128
		if hasWindowsPorts(openPorts) {
			return "Microsoft Windows"
		}
		return "Windows-like"
	case fp.TTL <= 255:
		// Some network devices use TTL 255
		return "Network device (Router/Switch)"
	}
	
	// Fallback to port-based detection
	return guessOSByPorts(openPorts)
}

func hasLinuxPorts(ports []int) bool {
	linuxPorts := []int{22, 80, 443, 25, 53, 110, 143, 993, 995}
	for _, port := range ports {
		for _, lport := range linuxPorts {
			if port == lport {
				return true
			}
		}
	}
	return false
}

func hasWindowsPorts(ports []int) bool {
	windowsPorts := []int{135, 139, 445, 3389, 1433, 1521}
	for _, port := range ports {
		for _, wport := range windowsPorts {
			if port == wport {
				return true
			}
		}
	}
	return false
}

func guessOSByPorts(ports []int) string {
	hasSSH := contains(ports, 22)
	hasHTTP := contains(ports, 80) || contains(ports, 8080)
	hasHTTPS := contains(ports, 443)
	hasSMB := contains(ports, 445) || contains(ports, 139)
	hasRDP := contains(ports, 3389)
	hasCUPS := contains(ports, 631)
	
	// Detect current OS if scanning localhost
	if hasCUPS {
		currentOS := runtime.GOOS
		switch currentOS {
		case "linux":
			return "Linux (detected via CUPS service)"
		case "darwin":
			return "macOS (detected via CUPS service)"
		case "windows":
			return "Windows (detected via runtime)"
		default:
			return "Unix-like with CUPS printing service"
		}
	}

	if hasRDP || hasSMB {
		return "Microsoft Windows"
	}
	if hasSSH && (hasHTTP || hasHTTPS) {
		return "Linux Server"
	}
	if hasSSH {
		return "Unix-like (Linux/BSD/macOS)"
	}
	if hasHTTP || hasHTTPS {
		return "Web Server (OS Unknown)"
	}

	return "Unknown"
}

func contains(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
