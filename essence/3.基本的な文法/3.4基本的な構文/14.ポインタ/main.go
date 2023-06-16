package main

import "fmt"

// ポインタ

func main() {
	v := 1
	fmt.Println(v)
	p := &v
	fmt.Println(p)
	*p = 2
	fmt.Println(v)
}
