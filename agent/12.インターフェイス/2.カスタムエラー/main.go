package main

import "fmt"

// interface
// custom error

// type error interface {
// 	Error() string
// }

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaseError() error {
	return &MyError{Message: "カスタムエラーが発生しました。", ErrCode: 444}
}

func main() {
	err := RaseError()
	fmt.Println(err.Error())
}
