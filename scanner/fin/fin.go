package fin

import (
	"fmt"
	"net"
	"syscall"
	"time"
	"unsafe"
)

type TCPHeader struct {
	Src    uint16
	Dst    uint16
	Seq    uint32
	Ack    uint32
	Flags  uint16
	Window uint16
	Csum   uint16
	Urg    uint16
}

func ScanPorts(host string, ports []int) {
	fmt.Printf("FIN scanning %s...\n\n", host)
	
	openFiltered := 0
	closed := 0
	
	for _, port := range ports {
		state := finScan(host, port)
		if state == "OPEN|FILTERED" {
			openFiltered++
		} else {
			closed++
			fmt.Printf("Port %d: %s\n", port, state)
		}
	}
	
	if openFiltered > 0 {
		fmt.Printf("\nNot shown: %d open|filtered ports\n", openFiltered)
	}
}

func finScan(host string, port int) string {
	// Try to create raw socket
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_TCP)
	if err != nil {
		// Fallback to connect scan if no raw socket privileges
		return connectFallback(host, port)
	}
	defer syscall.Close(fd)
	
	// Resolve host
	addr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		return "UNKNOWN"
	}
	
	// Create FIN packet
	packet := createFinPacket(uint16(port))
	
	// Send packet
	sockaddr := &syscall.SockaddrInet4{
		Port: port,
	}
	copy(sockaddr.Addr[:], addr.IP.To4())
	
	err = syscall.Sendto(fd, packet, 0, sockaddr)
	if err != nil {
		return "UNKNOWN"
	}
	
	// Wait for response
	buf := make([]byte, 1024)
	syscall.SetNonblock(fd, true)
	time.Sleep(100 * time.Millisecond)
	
	n, _, err := syscall.Recvfrom(fd, buf, 0)
	if err != nil || n == 0 {
		return "OPEN|FILTERED"
	}
	
	// Check for RST flag in response
	if len(buf) > 33 && buf[33]&0x04 != 0 {
		return "CLOSED"
	}
	
	return "OPEN|FILTERED"
}

func createFinPacket(dstPort uint16) []byte {
	header := TCPHeader{
		Src:    12345,
		Dst:    dstPort,
		Seq:    0,
		Ack:    0,
		Flags:  0x5001, // FIN flag + header length
		Window: 1024,
		Csum:   0,
		Urg:    0,
	}
	
	return (*(*[20]byte)(unsafe.Pointer(&header)))[:]
}

func connectFallback(host string, port int) string {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 200*time.Millisecond)
	if err != nil {
		return "CLOSED"
	}
	conn.Close()
	return "OPEN"
}