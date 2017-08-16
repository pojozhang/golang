package main

import (
	"fmt"
)

type Writer interface {
	Write(string)
}

type MailWriter struct {
}

type TextWriter struct {
}

func (w MailWriter) Write(text string) {
	fmt.Println("write " + text)
}

type WriteFunc func(string)

func (w WriteFunc) Write(text string) {
	w(text)
}

func main() {
	var writer Writer
	writer = new(MailWriter)
	fmt.Println(writer)
	writer.Write("hello")
	fmt.Printf("%T\n", writer) //打印接口值的动态类型

	//让函数值实现接口->起到适配器作用
	doWrite(WriteFunc(write), "func")
}

func doWrite(w Writer, s string) {
	w.Write(s)
}

func write(s string) {
	fmt.Println(s)
}
