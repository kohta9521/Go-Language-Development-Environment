package main

import "fmt"

// 定数

const Pi = 3.14

const (
	URL      = "https://google.com"
	SiteName = "Google"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

func main() {
	fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E)
}
