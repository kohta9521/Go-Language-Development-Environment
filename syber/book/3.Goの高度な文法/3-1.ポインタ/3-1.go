package main

import "fmt"

func main() {
	n := 123
	p := &n
	fmt.Println("number :", n)
	fmt.Println("number :", p)
	fmt.Println("number :", *p)
}
