package main

import "fmt"

// 繰り返し構文

func main() {
	// i := 0
	// for {
	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println("Looop")
	// }

	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// arr := [3]int{1, 2, 3}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// for _, v := range arr {
	// 	fmt.Println(v)
	// }

	// スライス　配列　可変長

	// sl := []string{"Python", "PHP", "Go"}
	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	m := map[string]int{"Apple": 100, "Banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
