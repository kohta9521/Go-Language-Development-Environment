package main

import "fmt"

// 型
// byte型

func main() {
	byteA := []byte{72, 73}
	// byte型はASCIIコードで表現される
	fmt.Println(byteA)
	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)
}