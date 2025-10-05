package main

import (
	"fmt"
	"math/rand"
	"time"
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

	if minNum >= maxNum {
		for minNum >= maxNum {
			fmt.Println("最小値よりも最大値を大きい整数にしてください。")
			fmt.Print("最小値: ")
			fmt.Scan(&minNum)
			fmt.Print("最大値: ")
			fmt.Scan(&maxNum)
		}
	}

	fmt.Println("OK!")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNum := r.Intn(maxNum-minNum+1) + minNum
	fmt.Println(randomNum)
}
