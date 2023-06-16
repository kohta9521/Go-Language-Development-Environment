package main

import (
	"fmt"
	"sort"
)

// map lesson

func main() {
	// m := make(map[string]int)
	// m["Jhon"] = 21
	// m["Bob"] = 18
	// m["Mark"] = 33

	// m := make(map[string]int, 10)
	// m["Jhon"] = 21
	// m["Bob"] = 18
	// m["Mark"] = 33

	m := map[string]int{
		"Jhon": 21,
		"Bob":  18,
		"mark": 22,
	}

	delete(m, "Bob")

	for k, v := range m {
		fmt.Printf("key: %v, value:  %v\n", k, v)
	}

	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("Key: %v, value: %v", k, m[k])
	}

	fmt.Println(m)
}
