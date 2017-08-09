package main

import (
	"fmt"
)

func main() {
	fmt.Println(f1())
	fmt.Println(f2())

	//函数变量
	f := f3
	fmt.Println(f(3))

	fmt.Println(f4()())
	fmt.Println(f4()())
	fmt.Println(f4()())

	ff := f4()
	fmt.Println(ff())
	fmt.Println(ff())
	fmt.Println(ff())

	var funcs []func()
	var r []int = []int{0, 1, 2}
	for _, i := range r {
		funcs = append(funcs, func() {
			fmt.Println(i) //i的取值由最后一次迭代决定
		})
	}

	for _, f := range funcs {
		f()
	}

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum([]int{1, 2, 3}...))//slice转成可变长参数
}

func f1() (name string, age int) {
	name = "name"
	age = 1
	return
}

func f2() (string, int) {
	return "name", 1
}

func f3(n int) int {
	return n * n
}

func f4() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func sum(numbers ...int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}
