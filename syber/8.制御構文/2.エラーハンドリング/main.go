package main

import (
	"fmt"
	"strconv"
)

// エラーハンドリング

func main() {
	var s string = "A"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", i)
}
