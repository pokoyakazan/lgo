package main

import "fmt"

func main() {
	// 3.1 配列
	{
		var x = [...]int{10, 20, 30}
		var y = [3]int{10, 20, 30}
		fmt.Println(x == y) // true
		x[0] = 10
		fmt.Println(x)
		fmt.Println("-----------------------------")
	}

	// スライス
	{
		var x = []int{10, 20, 30}
		fmt.Println(x)
	}
	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println(x)
		fmt.Println("-----------------------------")
	}
	// 3.2.6 スライスのスライス
	{
		x := []int{1, 2, 3, 4}
		y := x[:2]
		z := x[1:]
		d := x[1:3]
		e := x[:]
		fmt.Println("x: ", x) // [1 2 3 4]
		fmt.Println("y: ", y) // [1 2]
		fmt.Println("z: ", z) // [2 3 4]
		fmt.Println("d: ", d) // [2 3]
		fmt.Println("e: ", e) // [1 2 3 4]
		fmt.Println("-----------------------------")
	}
	// 複数のスライスによる要素の共有
	{
		x := []int{1, 2, 3, 4}
		y := x[:2]
		z := x[1:]
		x[1] = 20
		y[0] = 10
		z[1] = 30
		fmt.Println("x: ", x) // [10 20 30 4]
		fmt.Println("y: ", y) // [10 20]
		fmt.Println("z: ", z) // [20 30 4]
	}
	// 3.2.8 メモリを共有しないスライスのコピー
	{
		x := []int{1, 2, 3, 4}
		y := make([]int, 4)
		num := copy(y, x)   // num: コピーされた要素の数
		fmt.Println(y, num) // [1 2 3 4] 4
		fmt.Println("-----------------------------")
	}
	{
		x := []int{1, 2, 3, 4}
		y := make([]int, 2)
		num := copy(y, x)
		fmt.Println(y, num) // [1 2] 2
		fmt.Println("-----------------------------")
	}
	{
		x := []int{1, 2, 3, 4}
		y := make([]int, 2)
		copy(y, x[1:])
		fmt.Println(y) // [2 3]
		fmt.Println("-----------------------------")
	}
	{
		x := []int{1, 2, 3, 4}
		num := copy(x[:3], x[1:])
		fmt.Println(x, num) // [2 3 4 4] 3
		fmt.Println("-----------------------------")
	}
	// 3.3 文字列、rune、バイト
	{
		var a rune = 'x'
		var s1 string = string(a)
		var b byte = 'y'
		var s2 string = string(b)
		fmt.Println("s1: ", s1) // x
		fmt.Println("s2: ", s2) // y
	}
	{
		var s string = "Hello, 世界"
		var bs []byte = []byte(s)
		var rs []rune = []rune(s)
		fmt.Println("bs: ", bs)
		fmt.Println("rs: ", rs)
		/*
			bs:  [72 101 108 108 111 44 32 228 184 150 231 149 140]
			rs:  [72 101 108 108 111 44 32 19990 30028]
		*/
	}

}
