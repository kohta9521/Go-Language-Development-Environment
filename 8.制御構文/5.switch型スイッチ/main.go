package main

import "fmt"

// switch
// 型スイッチ


func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	anything("aaa")
	anything(1)

	var x interface{} = 3
	// 型アサーション
	i, isInt := x.(int)

	fmt.Println(i, isInt)
	fmt.Println(i + 2)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "x is string")
	} else {
		fmt.Println("I dont know")
	}

	// switch x.(type) {
	// case int:
	// 	fmt.Println("int")
	// case string:
	// 	fmt.Println("string")
	// case float64:
	// 	fmt.Println("float64")
	// case bool:
	// 	fmt.Println("bool")
	// default:
	// 	fmt.Println("I dont know")
	// }

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}
}