package main

import "fmt"

// 型
// 配列

func main() {
	// 要素を多言語とは違い増やしたりできない
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	// 初期値を設定
	// 文字列型
	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	// 暗黙的定義
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
	fmt.Printf("%T\n", arr3)

	// 要素数の省略
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	// 配列の操作
	fmt.Println(arr2[0])

	// 配列の値の更新
	arr2[2] = "C"
	fmt.Println(arr2)
	fmt.Printf("%T\n", arr2)

	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	// 配列の要素数を調べる
	fmt.Println(len(arr1))
}
