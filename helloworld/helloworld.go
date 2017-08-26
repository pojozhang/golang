package main

import (
	"./test"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
)

type COLOR string
type COLOR2 string

func main() {
	fmt.Println("hello world")
	printColor(COLOR("sss"))
	var red COLOR = "red"
	var red2 COLOR2 = "red"

	red += COLOR(red2)
	fmt.Println(red)
	fmt.Println(test.AKA)

	//switch语句，默认是break，需要显式声明fallthrough直接执行下一个case的语句，且最后一个case不能加fallthrough，否则编译错误
	d := 1
	switch d {
	case 2:
		fallthrough
	case 3:
	}
}

func printColor(color COLOR) {
	fmt.Println(color)
}

func init() {
	fmt.Println("init1")
	count := 0
	for range "中国" {
		count++
	}
	fmt.Println(count)
	fmt.Println(len("中国"))

}

func init() {
	fmt.Println("init2")
	fmt.Println(math.IsNaN(math.Sqrt(-1)))
}

var cwd string
var cwd2 string

func init() {
	fmt.Println("init3")
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalln("error:", err)
	}
}

func init() {
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString("中国")
	fmt.Println(buffer.String())
	ii := []int{1, 2}
	ii = []int{2, 3, 4}
	fmt.Println(ii)
	fmt.Println(len(ii))
}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func init() {
	symbol := []string{USD: "$", EUR: "aa"}
	fmt.Println(symbol)
	r := [...]int{15: 1} //定义一个长度为16的数组，且最后一个元素是1，其余元素是0
	r1 := r[2:4]
	fmt.Println(r1)
	fmt.Println(cap(r1))
	fmt.Println(r1[2:])
	ss(&r)
	fmt.Println(len(r))

	var nn []int = nil
	fmt.Println(len(nn))
}

func ss(b *[16]int) {

}
