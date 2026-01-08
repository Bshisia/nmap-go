package scanner

import (
	"net"
	"pentest-kit/scanner/aggressive"
	"pentest-kit/scanner/fin"
	"pentest-kit/scanner/null"
	"pentest-kit/scanner/os"
	"pentest-kit/scanner/service"
	"pentest-kit/scanner/syn"
	"pentest-kit/scanner/tcp"
	"pentest-kit/scanner/udp"
	"pentest-kit/scanner/xmas"
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

func UdpScan(host string, ports []int) {
	udp.ScanPorts(host, ports)
}

func FinScan(host string, ports []int) {
	fin.ScanPorts(host, ports)
}

func XmasScan(host string, ports []int) {
	xmas.ScanPorts(host, ports)
}

func NullScan(host string, ports []int) {
	null.ScanPorts(host, ports)
}

func OSDetection(host string, ports []int) {
	os.DetectOS(host, ports)
}

