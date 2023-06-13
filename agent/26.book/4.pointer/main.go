package main

import "fmt"

func main() {
	var x int32 = 10
	fmt.Println(" x =  ", x)
	fmt.Println("%x = ", &x)

	var p *int32
	p = &x
	fmt.Println(" p = ", p)
}
