package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := []int{10, 20, 30}
	fmt.Println(a)
	a = push(a, 1000)
	fmt.Println(a)
	a = pop(a)
	fmt.Println(a)
}

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}

func push(a []int, v int) []int {
	return append(a, v)
}

func pop(a []int) []int {
	return a[:len(a)-1]
}
