package main

import (
	"net"
	"fmt"
	"time"
	"log"
)

func main() {
	//打开连接:
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return
	}

	for {
		select {
		case <-time.After(1 * time.Second):
			_, err = conn.Write([]byte("hello"))
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
