package main

import "fmt"

type Person struct {
	firstName string
	age       int
}

func main() {
	tim := Person{"Tim", 25}
	person1 := &tim
	(*person1).age = 25
	person1.age = 53
	fmt.Println(person1)
}
