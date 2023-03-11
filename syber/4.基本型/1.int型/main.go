package main

import "fmt"

// 変数

func main() {
	var i int = 100

	var i2 int64 = 200
	fmt.Println(i + 50)
	fmt.Println(i2)

	fmt.Printf("%T\n", i2)

	fmt.Printf("%T\n", int32(i2))

	fmt.Println(int(i2) + i)
}
