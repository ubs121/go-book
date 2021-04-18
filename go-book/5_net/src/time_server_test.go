// DaytimeServer.go
package net

import (
	"net"
	"time"
)

func timeServer() {
	service := ":1200"
	listener, err := net.Listen("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // don't care about return value
		conn.Close()
		// we're finished with this client
	}
}
