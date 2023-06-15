package main

import "fmt"

// 定数宣言

func main() {
	const x = 1

	fmt.Println(x)

	const y = 2
	z := 1
	f := 1.2 + (y+2)*float64(z)
	fmt.Println(f)
}
