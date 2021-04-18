package net

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func echoServer() {
	service := ":1201"
	listener, err := net.Listen("tcp", service)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// run as a goroutine
		go handleEchoClient(conn)
	}
}

func handleEchoClient(conn net.Conn) {
	// close connection on exit
	defer conn.Close()
	var buf [512]byte

	// read upto 512 bytes
	n, err := conn.Read(buf[0:])
	if err != nil {
		return
	}

	fmt.Printf("Request: %s\n", string(buf[:]))

	// write back what it read
	_, err2 := conn.Write(buf[0:n])
	if err2 != nil {
		return
	}
}

func echoClient() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:1201")
	checkError(err)

	// холболт тогтоох
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("Hi"))
	checkError(err)

	// серверээс мэдээлэл унших
	result, err := io.ReadAll(conn)
	checkError(err)

	// авсан хариуг дэлгэцэнд хэвлэж харуулах
	fmt.Printf("Response: %s\n", string(result))
}

func TestEcho(t *testing.T) {

	go echoServer()

	echoClient()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
