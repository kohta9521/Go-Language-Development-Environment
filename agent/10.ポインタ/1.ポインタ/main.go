package main

import "fmt"

// pointer

func Double(i int) {
	i *= 2
}

// func Double2(i, *int) {
// 	*i = *i * 2
// }

func main() {
	var n int = 100
	fmt.Println(n)

	// アドレス表示
	fmt.Println(&n)

	Double(n)
	fmt.Println(n)

	// ポインタ型
	var p *int = &n
	fmt.Printf("%T\n", p)
	fmt.Println(p)

	*p = 300
	fmt.Println(n)
	n = 200
	fmt.Println(*p)

	fmt.Println(n)
}
