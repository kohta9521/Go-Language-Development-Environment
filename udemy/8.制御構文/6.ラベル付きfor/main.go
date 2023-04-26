package main

import "fmt"

// ラベル付き
// 条件分岐

func main() {
	// Loop:
	//
	//	for {
	//		for {
	//			for {
	//				fmt.Println("START")
	//				break Loop
	//			}
	//			fmt.Println("処理をしない")
	//		}
	//		fmt.Println("処理をしない")
	//	}
	//	fmt.Println("END")

	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			fmt.Println(i, j, i*j)
		}
	}
}
