package main

import (
	"fmt"
	"sync"
	"time"
)

func foo(wg *sync.WaitGroup, label string) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Println(label, i)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg, "A")
	wg.Add(1)
	go foo(&wg, "B")

	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("Done!!")
}
