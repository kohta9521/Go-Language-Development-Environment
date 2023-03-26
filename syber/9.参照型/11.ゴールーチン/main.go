package main

import "fmt"

// channel
// チャネルとゴルーチン

func recieve(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)

	// fmt.Println(<-ch1)

	ch2 := make(chan int)
	go recieve(ch2)

}
