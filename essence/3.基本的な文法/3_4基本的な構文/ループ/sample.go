package  main

import "fmt"

func main() [
	m := map[string]int {
		"John": 21
		"Bob": 18
		"Mark": 33
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Prinltn("key: %v, value : %v\n", k, v)
	}
]