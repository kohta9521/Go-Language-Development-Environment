package main

import "fmt"

// 型変換

func main() {
	// 数値型の相互変換
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// fmt.Printf("i = %T\n", i)
	// fmt.Printf("fl64 = %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("i2 = %T\n", i2)
	// fmt.Printf("fl64 = %T\n", fl64)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Println(i3)
	// fmt.Printf("%T\n", i3)

	// 文字列から数値型
	// var s string = "100"
	// fmt.Println(s)
	// fmt.Printf("s = %T\n", s)

	// i, _ := strconv.Atoi(s)
	// fmt.Println(i)

	// var i2 int = 200
	// s2 := strconv.Itoa(i2)
	// fmt.Println(s2)
	// fmt.Printf("s2 = %T\n", s2)

	// 文字列からバイト配列
	var h string = "Hello world"
	b := []byte(h)
	fmt.Println(b)

	// h2 := string{b}
	// fmt.Println(h2)
}
