package main

import (
	"fmt"
	"net"
	"testing"
)

func BenchmarkSpamTCP(b *testing.B) {
	conn, err := net.Dial("tcp", "localhost:7000")
	if err != nil {
		fmt.Println(err)
		panic("error establishing tcp connection")
	}
	defer conn.Close()

	spam(conn)
}

func BenchmarkSpamUDP(b *testing.B) {
	conn, err := net.Dial("udp", "localhost:7001")
	if err != nil {
		panic("error establishing udp connection")
	}
	defer conn.Close()

	spam(conn)
}
