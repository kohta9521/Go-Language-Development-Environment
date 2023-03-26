package main

import "fmt"

type Person struct {
	firstName string
	age       int
}

func main() {
	bob := Person{"Bob", 20}
	fmt.Println(bob.firstName, bob.age)
}
