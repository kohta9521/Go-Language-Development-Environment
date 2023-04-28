package main

import "fmt"

// スライス
// 可変長引数

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	fmt.Println(sum())

	// slice
	sl := []int{1, 2, 3}
	fmt.Println(sum(sl...))
}
