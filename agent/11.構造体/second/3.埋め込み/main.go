package main

import "fmt"

type T struct {
	User
}
type User struct {
	Name string
	Age  int
}

// func (u *User) SetName() {
// 	u.Nama = "A"
// }

func main() {
	t := T{User: User{Name: "user1", Age: 10}}
	fmt.Println(t)

	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	fmt.Println(t.User.Age)

}
