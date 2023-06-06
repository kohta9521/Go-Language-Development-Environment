package main

import "fmt"

// 型
// byte

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// byte型を文字列へ変換
	fmt.Println(string(byteA))

	c := []byte("Hi")
	fmt.Println(c)

	fmt.Println(string(c))
}
