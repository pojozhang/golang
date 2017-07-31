package main

import (
	"fmt"
	"strconv"
)

func main() {
	args := make(map[string]int)
	args["key"] = 1
	fmt.Println(args)

	args = map[string]int{
		"key": 1,
	}
	args["key2"] = 2
	fmt.Println(args)
	fmt.Println(args["key2"])
	fmt.Println(args["key3"])
	delete(args, "key4")

	for key, value := range args {
		fmt.Println(key + ":" + strconv.Itoa(value))
	}
}
