package main

import "fmt"

// マップ

func main() {
	var m = map[string]int{"A": 100, "B": 200}

	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}

	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}

	fmt.Println(m3)
}
