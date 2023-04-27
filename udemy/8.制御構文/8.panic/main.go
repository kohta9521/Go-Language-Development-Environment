package main

import "fmt"

// panic & recover
// 制御構文

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	// fmt.Println("Start")
}
