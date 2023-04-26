package main

import "fmt"

// switch型
// 条件分岐

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	anything("aaa")
	anything(1)

	var x interface{} = 3
	i := x.(int)
	fmt.Println(i + 2)
	fmt.Printf("%T\n", i)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)
	// fmt.Printf("%T\n", f)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I dont know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I dont know")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}
}
