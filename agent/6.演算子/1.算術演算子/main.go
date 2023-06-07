package main

import "fmt"

// 算術演算子

func main() {
	fmt.Println(1 + 2)
	fmt.Println(5 - 1)
	fmt.Println(6 * 4)
	fmt.Println(24 / 6)
	fmt.Println(53 % 3)

	fmt.Println("abc" + "def")

	// 算術演算子と代入
	n := 0
	n += 4
	fmt.Println(n)

	n++
	fmt.Println(n)
}
