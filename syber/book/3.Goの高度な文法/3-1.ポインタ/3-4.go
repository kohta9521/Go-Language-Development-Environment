package main

import "fmt"

func main() {
	n := 123
	fmt.Printf("value : %d.\n", n)

	cange1(n)
	fmt.Printf("value : %d.\n", n)

	cahnge2(&n)
	fmt.Print("value : %d.\n", n)
}

func cange1(n int) {
	n *= 2
}

func cahnge2(n *int) {
	*n *= 2
}
