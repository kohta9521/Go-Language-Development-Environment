package main

import "fmt"

// iota 列挙

const (
	Apple = iota + iota
	Orange
	Banana = iota + 3
)

// 型指定
// type Fruit int
// type Animal int

func main() {
	fmt.Println(Apple)
	fmt.Println(Orange)
	fmt.Println(Banana)

	// const (
	// 	Apple = iota
	// 	Orange
	// 	Banana
	// )

	// fmt.Println(Apple)
	// fmt.Println(Orange)
	// fmt.Println(Banana)
}
