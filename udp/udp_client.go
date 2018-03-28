package main

import (
	"net"
	"fmt"
	"time"
)

func main() {
	ip := net.ParseIP("127.0.0.1")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	fmt.Printf("<%s>\n", conn.RemoteAddr())

	for {
		select {
		case <-time.After(1 * time.Second):
			_, err = conn.Write([]byte("hello"))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
