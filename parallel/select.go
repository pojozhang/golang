package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("after 1 seconds")
	}
	fmt.Println("fire")

	ch := make(chan int, 1)
	for x := 0; x < 10; x++ {
		select {
		case ch <- x:
		case n := <-ch:
			fmt.Println(n)//print 0 2 4 6 8
		}
	}
}
