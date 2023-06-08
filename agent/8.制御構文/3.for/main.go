package main

import "fmt"

// if
// 繰り返し構文

func main() {
	// 	i := 0
	// 	for {
	// 		i++
	// 		if i == 3 {
	// 			break
	// 		}
	// 		fmt.Println("Loooop")
	// 	}

	// 条件付きfor
	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// 古典的for
	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for文で配列の中身を取り出す
	// arr := [3]int{1, 2, 3}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// 範囲式for
	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }

	// sl := []string{"Python", "PHP", "go"}
	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	// map

	m := map[string]int{"Apple": 100, "Banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
