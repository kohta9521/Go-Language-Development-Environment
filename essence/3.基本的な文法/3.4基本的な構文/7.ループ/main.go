package main

import "fmt"

// ループ

func main() {
	// c言語由来
	for i := 0; i < 5; i++ {
		fmt.Println("Looooop")
	}

	// label付き繰り返し構文
	// finished:
	// 	for a := 0; a < 100; a++ {
	// 		for j := 0; j < 100; j++ {
	// 			if check(a, j) {
	// 				break finished
	// 			}
	// 		}
	// 	}
}
