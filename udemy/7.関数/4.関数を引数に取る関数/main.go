package main

import "fmt"

// 関数
// 関数を引数に取る関数

func callFunction(f func()) {
	f()
}

func main() {
	callFunction(func() {
		fmt.Println("sample function")
	})
}
