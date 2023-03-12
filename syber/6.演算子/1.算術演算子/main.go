package main

import "fmt"

func main() {

	// 算術演算子
	fmt.Println(1 + 2)
	fmt.Println(5 - 1)
	fmt.Println(5 * 4)
	fmt.Println(60 / 3)
	fmt.Println(9 % 4)

	fmt.Println("ABC" + "DEF")

	n := 0
	// 省略技法
	n += 4
	fmt.Println(n)
	// インクリメント デクリメント
	n++
	fmt.Println(n)
	n--
	fmt.Println(n)

}
