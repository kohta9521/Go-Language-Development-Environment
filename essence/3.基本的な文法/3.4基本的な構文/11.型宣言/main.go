package main

import "fmt"

// type 型宣言

type MyString string

func main() {
	var m MyString
	m = "foo"
	fmt.Println(m)

	fmt.Printf("type: %v\n", m)
}
