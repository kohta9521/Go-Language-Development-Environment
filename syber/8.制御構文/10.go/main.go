package main

import (
	"fmt"
	"time"
)

// go gorutin

func sub() {
	for {
		fmt.Println("sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}
