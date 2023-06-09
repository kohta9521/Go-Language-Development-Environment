package main

import "fmt"

// 構造体

type User struct {
	Name string
	Age  int
	// X, Y int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	// fmt.Println(user1)

	// 明示的な変数定義
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	// 暗黙的
	user2 := User{}
	user2.Name = "user2"
	fmt.Println(user2)

	// 初期値を持たせて定義
	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	// userのnameなどを定義しない場合
	user4 := User{"user4", 40}
	fmt.Println(user4)

	// フィールドの順番に代入シアンければならない
	// user5 := User{30, "user5"}
	// fmt.Println(user5)

	// 片方のフィールドしか代入しない場合
	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User)
	fmt.Println(user7)

	user8 := &User{}
	fmt.Println(user8)

	// 関数を使用する場合
	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)

}
