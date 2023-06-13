package main

import "fmt"

// スライス

func printInfo(a []int) {
	s := fmt.Sprintf("%v", a)
	fmt.Printf("%-22s %3d\n", s, len(a), cap(a))
}

func main() {
	// a := []int{1, 2, 3}
	// fmt.Println(a)
	// fmt.Println(a[0])

	// スライスの長さとキャパシティについて
	a := []int{}

	// ヘッダの表示
	fmt.Printf("%-22s len cap\n", "values")
	for i := 0; i < 10; i++ {
		a = append(a, i+1)
		printInfo(a)
	}

	
}
