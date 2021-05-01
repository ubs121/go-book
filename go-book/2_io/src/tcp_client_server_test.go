package io

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func tcpServer() {
	service := ":8088"
	listener, err := net.Listen("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// клиентэд хариу илгээх
		conn.Write([]byte("Амжилттай холбогдлоо. Таны хаяг: " + conn.RemoteAddr().String() + "\n"))
		conn.Close()
	}
}

func tcpClient() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:8088")
	checkError(err)

	// холболт тогтоох
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	// серверээс мэдээлэл унших
	result, err := io.ReadAll(conn)
	checkError(err)

	// авсан хариуг дэлгэцэнд хэвлэж харуулах
	fmt.Println(string(result))
}

func TestTCP(t *testing.T) {

	go tcpServer()

	tcpClient()
}
