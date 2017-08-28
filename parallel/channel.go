package main

import (
	"fmt"
	"time"
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

	ch3 := make(chan int)
	//对于无缓冲通道，必须发送者和接受者成对出现，否则产生死锁，宕机，此处的go关键字不能省略
	go send(ch3)
	receive(ch3)

	//对于缓冲通道，可以同步写入，在队列未满的情况下不会产生死锁
	ch4 := make(chan int, 3)
	ch4 <- 1
	ch4 <- 2
	ch4 <- 3
	//如果在队列满的情况下发送消息，那么会报死锁错误
	//ch4 <- 4
	<-ch4
	ch4 <- 4

	//无缓冲通道是接近同步消费的，发送者需要有别的goroutine接手了才会继续执行，但不代表等接手的goroutine执行完后才继续执行
	ch5 := make(chan int)
	go func() {
		for n := 0; n < 10; n++ {
			fmt.Printf("send %d\n", n)
			ch5 <- n
		}
		close(ch5) //如果没有close也会产生死锁，因为已经没有发送者往通道里写消息
	}()

	for n := range ch5 {
		fmt.Printf("get %d\n", n)
		time.Sleep(300 * time.Millisecond)
	}

	//缓冲通道等同于异步消费，如果队列满了，发送方会阻塞，如果队列为空，消费方会阻塞
	ch6 := make(chan int, 10)
	go func() {
		for n := 0; n < cap(ch6); n++ {
			fmt.Printf("send %d\n", n)
			ch6 <- n
		}
		close(ch6)
	}()

	for n := range ch6 {
		fmt.Printf("get %d\n", n)
		time.Sleep(300 * time.Millisecond)
	}

	ch7 := make(chan struct{})
	go func() {
		time.Sleep(4 * time.Second)
		close(ch7)
	}()
	fmt.Println("wait")
	<-ch7//close后收到struct{}类型的零值，ok=false
	fmt.Println("done")
}

//只能发射的通道
func send(out chan<- int) {
	for n := 0; n < 10; n++ {
		out <- n
	}
	close(out)
}

//只能接收的通道
func receive(in <-chan int) {
	for n := range in {
		fmt.Printf("receive: %d\n", n)
	}
}
