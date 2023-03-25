package main

import "fmt"

// マップ

func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"C": 300, "D": 400}
	fmt.Println(m2)

	m3 := map[int]string{
		2: "b",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "Japan"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3])

	s, ok := m4[1]
	if !ok {
		fmt.Println("Erroe")
	}
	fmt.Println(s, ok)

	m4[2] = "US"
	fmt.Println(m4)

	// delete関数

	delete(m4, 2)
	fmt.Println(m4)

}
