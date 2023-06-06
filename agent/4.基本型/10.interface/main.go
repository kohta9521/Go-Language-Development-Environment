package main

import "fmt"

// 型
// interface

func main() {
	// 全ての型と互換性がある
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = "sample code"
	fmt.Println(x)

	fmt.Printf("%T\n", x)

	x = 3.14
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// 演算の対象としては利用不可能
	// 初期値としてnilを持っている
	// x = 2
	// fmt.Println(x + 3)

}
