package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(3))
	
	type company struct{
		name string `k:"v"`
		address string `k:"v"`
	}

	c:=company{}
	fmt.Println(reflect.ValueOf(c).Type().Field(0).Name)
	fmt.Println(reflect.ValueOf(c).Type().Field(0).Tag)

	n:=2
	reflect.ValueOf(&n).Elem().SetInt(3)
	fmt.Println(n)
}
