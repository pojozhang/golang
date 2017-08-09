package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"` //omitempty=> ignore in json if empty
}

func main() {

	books := []Book{
		{ID: 1, Name: "bookA"},
		{ID: 2, Name: "bookB"},
		{ID: 3, Name: "bookC"},
		{ID: 4},
	}
	if data, error := json.Marshal(books); error == nil {
		fmt.Printf("%s\n", data)
	}
	if data, error := json.MarshalIndent(books, "x", "  "); error == nil {
		fmt.Printf("%s\n", data)
	}

	var bks []Book
	fmt.Println(bks)
	data, _ := json.Marshal(books)
	json.Unmarshal(data, &bks)
	fmt.Println(bks)
}
