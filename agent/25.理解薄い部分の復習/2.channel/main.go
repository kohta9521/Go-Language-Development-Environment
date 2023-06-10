package main

import (
	"fmt"
	"sync"
)

func recv(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	s := <-ch
	fmt.Println("[R]", s)
}

func send(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[S] Sending...")
	ch <- "Hello"
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// チャネルの作成
	ch := make(chan string)
	// チャネルから文字列を受け取る側
	go recv(ch, &wg)
	// チャネルに文字列を送る側
	go send(ch, &wg)
	// 両方の goroutin が終わるのを待つ
	wg.Wait()

}
