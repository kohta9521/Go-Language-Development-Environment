package main

import "fmt"

// Mydata is structure.
type Mydata struct {
	Name string
	Data []int
}

func main() {
	kohta := Mydata{"Kohta", []int{10, 20, 30}}
	hanako := Mydata{
		Name: "Hanako",
		Data: []int{90, 80, 70},
	}

	fmt.Println(kohta)
	fmt.Println(hanako)
}
