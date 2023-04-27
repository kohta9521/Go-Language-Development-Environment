package main

import "fmt"

// スライス
// 宣言・操作

func main() {
	var sl []int
	fmt.Println(sl)

	// 明示的な宣言
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	// 暗黙的な宣言
	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	// make関数を使用
	sl4 := make([]int, 5)
	fmt.Println(sl4)

	// 要素の代入
	sl2[0] = 1000
	fmt.Println(sl2)

	// 値の取り出し
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)

	fmt.Println(sl5[0])
	fmt.Println(sl5[2:4])
	fmt.Println(sl5[:2])
}
