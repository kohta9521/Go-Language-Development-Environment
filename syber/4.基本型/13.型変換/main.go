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
	// fmt.Printf("%T\n", i)
	// fmt.Printf("%T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("%T\n", i2)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Printf("%T\n", i3)
	// fmt.Println(i3)

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	i, _ := strconv.Atoi("A")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {

	// }
	fmt.Println(i)
	fmt.Printf("s = %T\n", i)

	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	var h string = "Hello world"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)

}
