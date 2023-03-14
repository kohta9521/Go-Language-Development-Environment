package main

import "fmt"

// ラベル付き for

func main() {
	for {
		for {
			for {
				fmt.Println("START")
				break
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("処理をしない")
	}
	fmt.Println("END")
}
