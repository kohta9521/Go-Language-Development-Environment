package main

import "fmt"

// 関数

func Plus(x int, y int) int {
	return x + y
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)
}
