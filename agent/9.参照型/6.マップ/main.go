package main

import "fmt"

// マップ

func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	// 暗黙的宣言
	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	// make関数を使用
	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "Japan"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"])

	fmt.Println(m4[2])

	fmt.Println(m4[3])

	// error handlring
	s, ok := m4[4]
	fmt.Println(s, ok)
	if !ok {
		fmt.Println("this is not true number")
	}

	m4[2] = "US"
	fmt.Println(m4)

	m4[3] = "China"
	fmt.Println(m4)

	// delete関数の利用
	delete(m4, 3)
	fmt.Println(m4)

	// len関数の利用
	fmt.Println(len(m4))
}
