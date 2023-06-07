package main

import "fmt"

// 定数
// global部分に書くことが多い
const Pi = 3.14
const (
	URL      = "http://www.google.com"
	SiteName = "googe.com"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	// Pi = 3
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(c0, c1, c2)
}
