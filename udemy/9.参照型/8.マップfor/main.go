package main

import "fmt"

// map
// for

func main() {
	m := map[string]int{
		"apple":  100,
		"banana": 200,
	}

	// for
	for _, v := range m {
		fmt.Println(v)
	}
}
