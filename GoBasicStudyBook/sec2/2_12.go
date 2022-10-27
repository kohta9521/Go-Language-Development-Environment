package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := Input("type a numer")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Print("1から" + x + "の偶数の合計は")
	} else {
		return
	}
	t := 0
	c := 0
	for {
		c++
		if c%2 == 0 {
			continue
		}
		if c > n {
			break
		}
		t += c
	}
	fmt.Println(t, "です。")
}

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
