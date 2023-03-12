package main

import "fmt"

// switch 型
func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	anything("aaa")
	anything(1)
}
