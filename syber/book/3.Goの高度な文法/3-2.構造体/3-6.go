package main

import "fmt"

var mydata struct {
	Name string
	Data []int
}

func main() {
	mydata.Name = "Kohta"
	mydata.Data = []int{10, 20, 30}
	fmt.Println(mydata)
}
