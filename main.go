package main

import (
	"crypto/md5"
	"fmt"
	"net"
	"strconv"
)

func main() {}

func spam(conn net.Conn) {
	for i := 0; i < 100000; i++ {
		h := md5.Sum([]byte(strconv.Itoa(i)))
		fmt.Fprintf(conn, "%x", h)
	}
}
