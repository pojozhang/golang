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
