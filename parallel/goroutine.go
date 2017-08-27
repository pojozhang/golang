package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Println("print")
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(3 * time.Second)
}
