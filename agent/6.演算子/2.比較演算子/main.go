package main

import "fmt"

// 比較演算子

func main() {
	fmt.Println(1 == 1)

	fmt.Println(1 == 2)

	fmt.Println(4 <= 8)

	fmt.Println(8 <= 1)

	fmt.Println(1 < 8)

	// 論理値をそのまま比較
	fmt.Println(true == false)
	fmt.Println(true != false)
}
