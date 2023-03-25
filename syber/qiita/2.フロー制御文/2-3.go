package main

import "fmt"

func condition(arg string) string {
	if v := "Go"; arg == v {
		return "This is Golang"
	} else {
		return "This is not Goland"
	}
}

func main() {
	fmt.Println(condition("Swift"))
}
