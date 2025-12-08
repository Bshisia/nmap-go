package syn

import (
	"fmt"
	"net"
	"syscall"
	"time"
	"unsafe"
)

func ScanPorts(host string, ports []int) {
	fmt.Printf("SYN scanning %s...\n", host)
	
	for _, port := range ports {
		if probe(host, port) {
			fmt.Printf("Port %d: OPEN\n", port)
		}
	}
}

func probe(host string, port int) bool {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", addr, 500*time.Millisecond)
	if err == nil {
		defer conn.Close()
		tcpConn := conn.(*net.TCPConn)
		file, _ := tcpConn.File()
		fd := int(file.Fd())
		
		var info syscall.TCPInfo
		infoLen := uint32(unsafe.Sizeof(info))
		_, _, errno := syscall.Syscall6(syscall.SYS_GETSOCKOPT, uintptr(fd), syscall.SOL_TCP, syscall.TCP_INFO, uintptr(unsafe.Pointer(&info)), uintptr(unsafe.Pointer(&infoLen)), 0)
		
		file.Close()
		return errno == 0
	}
	return false
}