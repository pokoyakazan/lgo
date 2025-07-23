package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	// 4.2 if
	{
		// if文の変数のスコープ
		// nはelseブロックまで有効
		if n := rand.Intn(10); n == 0 {
			fmt.Println("n is ", 0)
		} else if n > 5 {
			fmt.Println("n is greater than 5: ", n)
		} else {
			fmt.Println("n is less than or equal to 5: ", n)
		}
		// fmt.Println("n is ", n) -> undefined: n
		fmt.Println("-----------------------------")
	}
	// 4.3 for
	{
		// シンプル
		for i := 0; i < 10; i++ {
			fmt.Println("i is ", i)
		}
		fmt.Println("---")
		// 条件のみ
		i := 1
		for i < 20 {
			fmt.Println("i is ", i)
			i *= 2
		}
		fmt.Println("---")
		// 無限ループ
		i = 1
		for {
			fmt.Println("i is ", i)
			i *= 2
			if i > 20 {
				break
			}
		}
		fmt.Println("---")
		// for-rangeループ
		evenVals := []int{0, 2, 4, 6, 8}
		for i, v := range evenVals {
			fmt.Println(i, v)
		}
		for _, v := range evenVals {
			fmt.Println(v)
		}
		fmt.Println("---")
		uniqueNames := map[string]bool{"hanako": true, "sakura": true}
		for k := range uniqueNames {
			fmt.Println(k)
		}
		fmt.Println("---")
		// マップは順序が保証されていないので、順番はランダム
		m := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}
		for i := 0; i < 3; i++ {
			fmt.Println("Loop:", i)
			for k, v := range m {
				fmt.Println(k, v) // 実行の度変わる
			}
		}
		fmt.Println("---")
		// for-rangeの値はコピー(値渡し)
		evenVals = []int{0, 2, 4, 6, 8}
		for _, v := range evenVals {
			v *= 2
		}
		fmt.Println(evenVals) // [0 2 4 6 8]
		fmt.Println("---")
		// ラベル
		samples := []string{"hello", "apple", "これは日本語"}
	outer:
		for _, s := range samples {
			for i, r := range s {
				fmt.Println(i, r, string(r))
				if r == 'l' || r == 'は' {
					continue outer
				}
			}
			fmt.Println()
		}
		fmt.Println("-----------------------------")
	}
	// 4.4 switch -> 全飛ばし、必要になったら見る
	// 4.6 goto: 基本使わないが使ったほうが良い一例
	{
		rand.Seed(time.Now().Unix())
		a := rand.Intn(10)
		for a < 100 {
			fmt.Println(a)
			if a%5 == 0 {
				goto done
			}
			a = a*2 + 1
		}
		fmt.Println("ループが通常終了")
	done:
		fmt.Println("gotoで終了")
		fmt.Println(a)
		fmt.Println("-----------------------------")
	}
}
