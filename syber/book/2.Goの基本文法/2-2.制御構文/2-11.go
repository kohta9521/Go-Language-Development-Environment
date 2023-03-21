package main

import (
	"fmt"
	"hello"
	"strconv"
)

func main() {
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)

	if err == nil {
		fmt.Print("1から" + x + "までの合計の値は、")
	} else {
		return
	}

	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}

	fmt.Println(t, "です。")
}
