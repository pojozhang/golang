package main

import (
	"fmt"
	"net/http"
)

func main() {
	//method1
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", handle)
	// http.ListenAndServe("localhost:8000", mux)

	//method2 默认使用DefaultServeMux
	http.HandleFunc("/", handle)
	http.ListenAndServe("localhost:8000", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello")
}
