package main

import "fmt"

type Person struct {
	firstName string
	age       int
}

func newPerson(firstName string, age int) *Person {
	person := new(Person)
	person.firstName = firstName
	person.age = age
	return person
}

func main() {
	var jen *Person = newPerson("Jennifer", 40)
	fmt.Println(jen.firstName, jen.age)
}
