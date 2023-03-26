package main

import "fmt"

type Person struct {
	firstName string
	age       int
}

func main() {
	sam := Person{age: 15, firstName: "sam"}
	fmt.Println(sam.firstName, sam.age)
}
