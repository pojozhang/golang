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

func main() {
	var writer Writer
	writer = new(MailWriter)
	fmt.Println(writer)
	writer.Write("hello")
}
