package main

import "fmt"

// 条件分岐

func main() {

	var i = 3

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("other number")
	}
}
