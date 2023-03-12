package main

import "fmt"

// 定数

const Pi = 3.14

const (
	URL      = "http://google.com"
	SiteName = "google"
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
	// 連続する整数を出力
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	// Pi = 3
	// fmt.Pirntln(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)
	fmt.Println(c0, c1, c2)
}
