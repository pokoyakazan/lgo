package main

import (
	"fmt"
	"time"
)

// 7.1 型の定義
type Score int                    // intを基底型としてScore型を定義
type Converter func(string) Score // stringを引数にScoreを返す関数型を定義
type TeamScores map[string]Score  // チーム名をキー、Scoreを値とするマップ型を定義

// 7.2 メソッド
type Person struct {
	LastName  string
	FirstName string
	Age       int
}

func (p Person) String() string { // p Person型のレシーバを持つ
	return fmt.Sprintf("%s %s (%d)", p.FirstName, p.LastName, p.Age)
}

// 7.2.1 ポインタ型レシーバと値型レシーバ
type Counter struct {
	total       int
	lastUpdated time.Time
}

// 以下のようにポインタレシーバと値レシーバをそれぞれ用意するのは一般的にはNG
func (c *Counter) Increment() { // ポインタレシーバ(cはポインタ)
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) StringC() int { // 値型レシーバ(cにはコピーが渡される)
	return fmt.Sprintf("合計: %d, 更新: %v", c.total, c.lastUpdated)
}

func main() {
	// 7.2 メソッド
	{
		// Person型のインスタンスを作成
		p := Person{LastName: "山田", FirstName: "太郎", Age: 30}
		// Stringメソッドを呼び出して文字列を出力
		f := p.String()
		fmt.Println(f) // 出力: 太郎 山田 (30)
	}
	// 7.2.1 ポインタ型レシーバと値型レシーバ
	{
		var c Counter
		fmt.Plintln(c.lastUpdated.StringC())

	}

}
