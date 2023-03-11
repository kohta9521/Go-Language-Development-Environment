package main

import "fmt"

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	//可変長配列 slice

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)
}
