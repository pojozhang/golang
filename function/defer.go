package main

import (
	"log"
	"time"
)

func main() {
	defer trace()()
	log.Println("run")
	time.Sleep(3 * time.Second)
}

func trace() func() {
	start := time.Now()
	log.Println(start)
	return func() {
		log.Println(time.Since(start))
	}
}

func init() {
	sum(1, 2, 3)
}

func sum(values ...int) (result int) {
	defer log.Printf("result: %d\n", result)
	for _, v := range values {
		result += v
	}
	return
}

func init(){
	defer log.Println("1")
	defer log.Println("2")
	defer log.Println("3")
	panic("crash")
}
