package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	slice := IntSlice{3, 8, 2}
	fmt.Println(slice)
	sort.Sort(slice)
	fmt.Println(slice)
}
