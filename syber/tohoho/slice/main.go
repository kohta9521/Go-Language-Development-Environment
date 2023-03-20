package main

import "fmt"

func main() {
	a1 := []string{}
	a1 = append(a1, "red")
	a1 = append(a1, "Green")
	a1 = append(a1, "blue")
	fmt.Println(a1[0], a1[1], a1[2])

	a := []int{}
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Println(len(a), cap(a))
	}

	// bufa := make([]byte, 0, 1024)
}
