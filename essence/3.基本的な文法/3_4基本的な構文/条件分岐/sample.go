package main

import (
	"fmt"
	"log"
)

func doSomething() {
	fmt.Println("doing something")
}

func main() {
	// x := 3
	// y := 3
	// if x == 3 && y == 3 {
	// 	doSomething()
	// }

	if user, err := userName(); err == nil {
		fmt.Println(user.Name)
	} else {
		log.Println(err)
	}
}
