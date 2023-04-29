package main

import "fmt"

// ポインタ

func Double(i int) {
	i = i * 2
}

func Double2(i *int) {
	*i = *i * 2
}

func main() {
	var n int = 100
	fmt.Println(n)

	fmt.Println(&n)

	Double(n)
	fmt.Println(n)

	// ポインタ型
	var p *int = &n
	fmt.Println(p)

	*p = 300
	fmt.Println(n)
	fmt.Println(&p)
}
