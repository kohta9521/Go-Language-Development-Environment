package main

import "fmt"

// switch
//  型によるSwitch

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	// 型アサーション
	anything("aaa")
	anything(1)

	var x interface{} = 3
	i, isInt := x.(int)
	fmt.Println(i+2, isInt)

	// f := x.(float64)
	// fmt.Println(f)

	// 型アサーションで停止させない方法
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	// 型アサーションで条件分岐
	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I dont know")
	}

	// 型によるスイッチ
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("i dont know")
	}

}
