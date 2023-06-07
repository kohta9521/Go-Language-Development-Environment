package main

import "fmt"

// 制御構文
// if

func main() {
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I dont know")
	}

	// 簡易文付き
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	//注意点 内部の変数のあたいが優先される
	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)
}
