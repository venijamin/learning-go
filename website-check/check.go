package main

import (
	"fmt"
	"net"
	"time"
)

func Check(domain string, port string) string {
	address := domain + ":" + port
	timeout := time.Duration(5 * time.Second)

	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, ERROR %v", address, err)
		return status
	}

	status = fmt.Sprintf("[UP] %v is reachable \n %v : %v", address, conn.LocalAddr(), conn.RemoteAddr())
	return status
}
