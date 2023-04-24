package main

import (
	"fmt"
	"strconv"
)

// 型変換

func main() {
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// fmt.Printf("i = %T\n", i)
	// fmt.Printf("fl64 = %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("%T\n", i2)

	// fl := 18.5
	// i3 := int(fl)
	// fmt.Println("i3 = %T\n", i3)
	// fmt.Println(i3)

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
