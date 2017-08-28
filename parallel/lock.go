package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock sync.Mutex
)

func main() {

	var sum int

	for n := 0; n < 100; n++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			sum++
			fmt.Println(sum)
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println(sum)
}
