package main

import "fmt"

// slice
// map

func main() {
	var m = map[string]int{"A": 100., "B": 200}
	fmt.Println(m)

	// 暗黙的宣言
	m2 := map[string]int{"C": 300, "D": 400}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	// make関数を利用
	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "Japan"
	m4[2] = "USA"
	fmt.Println(m4)

	// 値の取り出し
	fmt.Println(m["A"])
	fmt.Println(m4[2])

	fmt.Println(m4[3])

	_, ok := m4[4]
	if !ok {
		fmt.Println("error")
	}
	// fmt.Println(ok)

	// 削除
	delete(m4, 2)
	fmt.Println(m4)

	// len関数
	fmt.Println(len(m4))
}
