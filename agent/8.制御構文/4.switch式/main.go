package main

import "fmt"

// switch
// 式によるSwitch

func main() {
	// n := 40
	// switch n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// case 5, 6:
	// 	fmt.Println("5 or 6")
	// default:
	// 	fmt.Println("I dont know")
	// }

	// switch n := 2; n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("i dont know")
	// }

	n := 6
	switch {
	case n > 0 && n < 4:
		fmt.Println("0 < n < 4")
	}
}
