package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := Input("type data")
	ar := strings.Split(x, " ")
	t := 0
	for i := 0; i < len(ar); i++ {
		n, er := strconv.Atoi(ar[i])
		if er != nil {
			fmt.Println("ERROR!!")
		}
		t += n
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
