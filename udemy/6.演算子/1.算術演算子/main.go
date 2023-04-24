package main

import "fmt"

// 算術演算子

func main() {
	fmt.Println(1 + 2)
	fmt.Println(5 - 1)
	fmt.Println(5 * 4)
	fmt.Println(60 / 3)
	fmt.Println(9 % 4)

	// 文字列でも使用可能
	fmt.Println("ABC" + "DEF")

	// 算術演算子と代入
	a := 0
	b := 4
	fmt.Println(a * b)
}
