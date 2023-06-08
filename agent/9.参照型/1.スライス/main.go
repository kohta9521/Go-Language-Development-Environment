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

	// make関数を利用
	sl4 := make([]int, 5)
	fmt.Println(sl4)

	// 値の代入
	sl2[0] = 1000
	fmt.Println(sl2)

	// 値の取り出し
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5[2])

	// 範囲を持ったスライスの取り出し
	fmt.Println(sl5[2:4])

	fmt.Println(sl5[:2])

	fmt.Println(sl5[2:])

	fmt.Println(sl5[:])

	fmt.Println(sl5[len(sl5)-1])

	fmt.Println(sl5[1 : len(sl5)-1])
}
