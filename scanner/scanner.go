package scanner

import (
	"net"
	"pentest-kit/scanner/aggressive"
	"pentest-kit/scanner/service"
	"pentest-kit/scanner/syn"
	"pentest-kit/scanner/tcp"
)

func ScanPorts(host string, ports []int, serviceDetection bool) {
	tcp.ScanPorts(host, ports, serviceDetection)
}

func DetectService(conn net.Conn, port int) string {
	return service.DetectService(conn, port)
}

func SynScan(host string, ports []int) {
	syn.ScanPorts(host, ports)
}

func AggressiveScan(host string, ports []int) {
	aggressive.ScanPorts(host, ports)
}

