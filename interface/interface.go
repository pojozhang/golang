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

type PaperWriter struct {
}

func (w PaperWriter) Write(text string) {
	fmt.Println("write " + text)
}

type WriteFunc func(string)

func (w WriteFunc) Write(text string) {
	w(text)
}

func main() {
	var writer Writer = &MailWriter{}
	fmt.Println(writer)
	writer.Write("hello")
	fmt.Printf("%T\n", writer) //打印接口值的动态类型

	//让函数值实现接口->起到适配器作用
	doWrite(WriteFunc(write), "func")

	//类型断言
	w2 := writer.(Writer)
	w2.Write("w2")

	//类型断言，注意，此处判断的是实现了某个接口的某个具体类型，I.(T) 类型T必须实现了接口I，否则编译报错
	var w3 Writer = &MailWriter{}
	if f, ok := w3.(*PaperWriter); !ok { //ok为false，因为w3是*MailWriter类型
		fmt.Println(f)
		fmt.Println("not match")
	}

	formatValue(2)
	formatValue("2")
}

func doWrite(w Writer, s string) {
	w.Write(s)
}

func write(s string) {
	fmt.Println(s)
}

//类型分支
func formatValue(x interface{}) {
	//写法1
	if _, ok := x.(int); ok {
		fmt.Println("is int")
	}
	if _, ok := x.(string); ok {
		fmt.Println("is string")
	}

	//写法2，type是关键字，且类型分支不支持fallthrough
	switch x := x.(type) {
	case int:
		fmt.Println(x)//x是switch语句里创建的局部变量x，而不是外部变量
		fmt.Println("is int")
	case string:
		fmt.Println("is string")
	default:
		fmt.Println("is unknown")
	}
}
