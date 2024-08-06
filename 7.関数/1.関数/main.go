package main

import "fmt"

// 関数

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	println(i)
	i2, _ := Div(9, 2)
	fmt.Println(i2)

	i3 := Double(1000)
	fmt.Println(i3)

	Noreturn()
}