package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// x := Input("type a number")
	// n, err := strconv.Atoi(x)
	// p := float64(n)
	// fmt.Println(int(p * 1.1))
}

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
