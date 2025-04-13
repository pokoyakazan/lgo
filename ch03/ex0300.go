package main

import "fmt"

func main() {
	// 配列
	var x = [...]int{10, 20, 30}
	var y = [3]int{10, 20, 30}
	fmt.Println(x == y) // true
	x[0] = 10
	fmt.Println(len(x))

	// スライス
	var x = []int{10, 20, 30}

}
