package main

import (
	"fmt"
	"log"
)

func FindUser(name string) (*User, error) {
	user, err := findUserFromList(users, name)
	if err != nil {
		return nil, err
	}
	return user, err
}

func main() {
	user, err := FindUser("Bob")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.name)
}
