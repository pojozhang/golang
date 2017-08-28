package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for n := 0; n < 10; n++ {
			ch1 <- n
		}
		close(ch1)
	}()

	go func() {
		for {
			n, ok := <-ch1
			if ok {
				ch2 <- n * n
			} else {
				break
			}
		}
		close(ch2)
	}()

	//range方式接收通道里的消息，当通道被close后自动退出循环
	for n := range ch2 {
		fmt.Println(n)
	}
}
