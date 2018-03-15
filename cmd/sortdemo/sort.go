package main

import (
	"sort"
	"fmt"
)

func main() {
	// slice of int
	a := []int{3, 6, 2, 1, 9, 10, 8}

	sort.Ints(a)

	// “_”占位符，不使用，但是range的输出有
	for _, v := range a {
		fmt.Println(v)
	}
}
