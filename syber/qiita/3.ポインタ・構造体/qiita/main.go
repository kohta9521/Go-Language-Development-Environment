package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// ポインタ型の変数を宣言する
	// がポインタ変数
	// *Person がポインタ型
	var p *Person

	p = &Person{
		Name: "太郎",
		Age:  20,
	}

	fmt.Printf("変数pに格納されているアドレス : %p", p)
}
