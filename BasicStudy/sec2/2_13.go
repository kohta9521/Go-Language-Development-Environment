package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mian() {

	x := Input("type a number")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("ERROR!!")
	}
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total:", t)
	return

}

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
