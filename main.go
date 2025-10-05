package main

import (
	"fmt"
)

func main() {
	fmt.Println("数あてゲームを始めましょう。")
	fmt.Println("私が整数を思い浮かべるので、それを当ててください。")

	fmt.Println("ゲームを始める前に、数字の範囲を決めます。")
	fmt.Print("まずは、最小値を整数で入力してください: ")

	var minNum int
	fmt.Scan(&minNum)
	fmt.Println("最小値: ", minNum)

	fmt.Print("次に、最大値を整数で入力してください: ")
	var maxNum int
	fmt.Scan(&maxNum)
	fmt.Println("最大値: ", maxNum)
}
