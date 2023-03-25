package main

import "fmt"

type Person struct {
	firstName string
	age       int
}

func main() {
	// var mike Person
	// mike.firstName = "Mike"
	// mike.age = 20
	// fmt.Println(mike.firstName, mike.age)

	bob := Person{"Bob", 30}
	fmt.Println(bob.firstName, bob, age)
}
