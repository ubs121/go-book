// GetHeadInfo.go
package io

import (
	"fmt"
	"io"
	"net"
	"os"
	"testing"
)

func TestHeadInfo(t *testing.T) {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	//result, err := readFully(conn)
	result, err := io.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}
