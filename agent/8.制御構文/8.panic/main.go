package main

import "fmt"

// panic
// 例外処理
// プログラムの実行中にエラーが発生した場合に、プログラムを停止さ

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("rutime error")
	fmt.Println("Start")
}
