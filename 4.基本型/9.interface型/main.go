package main

import "fmt"

// 型
// interface型

func main() {
	var x interface{}
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 1
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 3.14
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = "Golang"
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}