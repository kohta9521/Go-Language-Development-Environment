package main

import "fmt"

func main() {
	lang := "Go"

	switch lang {
	case "Ruby":
		fmt.Println("This is Ruby")
	case "Go":
		fmt.Println("This is Go")
	default:
		fmt.Println("This is a programming Languag")
	}
}
