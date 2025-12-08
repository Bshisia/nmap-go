package tcp

import (
	"pentest-kit/scanner/service"
	"net"
)

func DetectService(conn net.Conn, port int) string {
	return service.DetectService(conn, port)
}