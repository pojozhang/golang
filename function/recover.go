package main

import (
	"fmt"
	"time"
)

func main() {

	type exp struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			fmt.Println("nil")
		case exp{}:
			fmt.Println("exp")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("sss")
	panic(exp{})
}
