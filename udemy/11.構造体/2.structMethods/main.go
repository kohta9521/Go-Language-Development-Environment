package main

import "fmt"

// struct
// method

type User struct {
	Name string
	Age  int
}

// method
func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) setName(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName()

	// user1.SetName("A")
	// user1.Sayname()
}
