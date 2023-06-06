package main

import "fmt"

// 文字列型

func main() {
	var s string = "Hello"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// 数値を文字列として扱う
	var s1 string = "300"
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)

	// 複数業の文字列
	fmt.Println(`test
		sample 
		code
	`)

	// 文字列の表現
	fmt.Println("\"sample dauble")

	// 文字列を取得する
	fmt.Println(string(s[0]))
	fmt.Println(s[1])
}
