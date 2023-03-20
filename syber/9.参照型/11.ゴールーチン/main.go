package main

// channel
// close

func main() {
	ch1 := make(chan int, 2)
	close(ch1)

	ch1 <- 1
}
