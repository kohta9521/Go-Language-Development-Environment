package main

import "fmt"

// 配列

func main() {
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"
	fmt.Println(fruits)

	// 初期化を同時に行う
	var fruits2 [3]string = [3]string{"Apple", "Banana", "Orange"}
	fmt.Println(fruits2)

	// 配列の比較
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	fmt.Println(a == b)
	b[1] = 5
	fmt.Println(a == b)

	// forループ
	c := [...]rune{'A', 'B', 'C'}
	fmt.Printf("Type: %T\n", c)
	for i, ch := range c {
		fmt.Printf(" c[%d] = %c\n", i, ch)
	}
}
