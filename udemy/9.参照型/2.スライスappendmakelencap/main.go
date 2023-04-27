package main

import "fmt"

// スライス
// append make len cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	// append
	// 追加する
	sl = append(sl, 300)
	fmt.Println(sl)

	sl = append(sl, 400, 500)
	fmt.Println(sl)

	// make
	// sliceを生成する
	sl2 := make([]int, 2)
	fmt.Println(sl2)
	fmt.Println(len(sl2))

	// slice cap
	fmt.Println(cap(sl2))

	sl3 := make([]int, 10)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl))

	// パフォーマンス面で考慮が必要
	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))
}
