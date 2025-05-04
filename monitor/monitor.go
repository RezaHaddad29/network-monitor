package monitor

import (
	"net"
	"time"
)

func CheckServer(address string, port string) bool {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(address, port), timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
