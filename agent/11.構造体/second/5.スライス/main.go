package main

// slice

type User struct {
	Name string
	Age  int
}

type Users []*User

func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}

	users := Users{}
	users = append(users, &user1) // *User
}
