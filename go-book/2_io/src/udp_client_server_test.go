package io

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func udpServer() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleUdpClient(conn)
	}
}

func handleUdpClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	fmt.Println(addr, " хаягаас хүсэлт ирлээ")
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}

func udpClient() {
	service := "localhost:1200"
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	_, err = conn.Write([]byte("хүсэлт"))
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)

	fmt.Println("Серверийн хариу", string(buf[0:n]))
}

func TestUDP(t *testing.T) {
	go udpServer()

	udpClient()
}
