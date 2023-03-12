package main

import "fmt"

// 条件分岐

func main() {
	a := 1
	if a == 2 {
		fmt.Println("Two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I dont know")
	}

	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)
}
