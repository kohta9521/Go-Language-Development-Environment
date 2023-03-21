package main

import (
	"fmt"
	"hello"
	"strconv"
	"strings"
)

func main() {
	x := hello.Input("input data")
	ar := strings.Split(x, " ")
	t := 0
	for _, v := range ar {
		n, err := strconv.Atoi(v)
		if err != nil {
			goto er
		}
		t += n
	}
	fmt.Println("total :", t)
	return

er:
	fmt.Println("ERROE")
}
