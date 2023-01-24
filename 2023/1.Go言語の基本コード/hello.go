package main

import (
	"fmt"
	"hello"
)

func main() {
	name := hello.Input("type a number")
	fmt.Println("Hello," + name + "!")
}
