package main

import "fmt"

// interface型

func main() {
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
}
