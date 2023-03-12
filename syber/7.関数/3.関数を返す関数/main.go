package main

import "fmt"

// 関数を返す関数

func Return() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func main() {
	f := Return()
	f()
}
