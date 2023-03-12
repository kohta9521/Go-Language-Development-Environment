package main

import "fmt"

// 繰り返し構文

func main() {
	i := 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("Looop")
	}
}
