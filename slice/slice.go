package main

import (
	"fmt"
)

func main() {
	appendElementsToSlice()
}

func appendElementsToSlice() {
	var slice []rune
	slice = append(slice, 1)
	slice = append(slice, 2, 3)
	slice = append(slice, slice...)
	fmt.Println(slice)
}
