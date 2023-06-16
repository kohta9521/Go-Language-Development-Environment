package main

import "fmt"

// 配列とスライス

func main() {

	// 長さを指定しているためこれは配列
	var a [4]int
	a[0] = 1
	fmt.Println(a)

	// 長さを指定していないためこれはスライスとなる
	var n []int
	fmt.Println(n)

	n = append(n, 4, 5, 6)
	fmt.Println(n)

	// make
	b := make([]int, 3)
	b[0] = 1
	b[1] = 2
	b[2] = 3

	fmt.Println(b)

	arr1 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(arr1)

	// スライスの2次元
	arr2 := [][]int{
		{1, 2, 3},
		{3, 4, 5},
	}
	fmt.Println(arr2)

	e := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		e = append(e, i)
	}
}
