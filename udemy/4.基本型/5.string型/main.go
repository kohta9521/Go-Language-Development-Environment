package main

import "fmt"

// 型
// 文字列型

func main() {
	var s string = "Hello Go Lang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	fmt.Println(`text
	text
	test
	`)

	fmt.Println("\"")

	fmt.Println(string(s[0]))
}
