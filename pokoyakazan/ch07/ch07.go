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

func (c Counter) StringC() string { // 値型レシーバ(cにはコピーが渡される)
	return fmt.Sprintf("合計: %d, 更新: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) { // 間違い
	c.total++ // cのコピーに対してIncrementが行われる
	c.lastUpdated = time.Now()
}

func doUpdateRight(c *Counter) { // 正しい
	c.total++ // cに対してIncrementが行われる
	c.lastUpdated = time.Now()
}

// 7.2.2 nilへの対応
type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(v int) *IntTree {
	if it == nil {
		return &IntTree{val: v} // nilの場合は新しいIntTreeを返す
	}
	if v < it.val {
		it.left = it.left.Insert(v)
	} else if v > it.val {
		it.right = it.right.Insert(v)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

// 7.2.3 メソッドは関数
type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
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
	fmt.Println("-----------------------------")
	// 7.2.1 ポインタ型レシーバと値型レシーバ
	{
		var c Counter
		fmt.Println(c)
		fmt.Println(c.StringC())
		c.Increment() // (&c).Increment()と書かなくても良い
		fmt.Println(c.StringC())
		doUpdateWrong(c)
		fmt.Println(c.StringC())
		doUpdateRight(&c)
		fmt.Println(c.StringC())
	}
	fmt.Println("-----------------------------")
	// 7.2.2 nilへの対応
	{
		var it *IntTree
		it = it.Insert(5) // nilの場合は新しいIntTreeを返す
		fmt.Println(*it)  // 5 <nil> <nil>}
		it = it.Insert(3)
		fmt.Println(*it)      // 5 0x14000142090 <nil>}
		fmt.Println(*it.left) // {3 <nil> <nil>}
		it = it.Insert(10)
		it = it.Insert(2)
		fmt.Println(it.Contains(2))  // true
		fmt.Println(it.Contains(12)) // false
		fmt.Println(*it)             // 5 0x14000142078 0x14000142090}
	}
	fmt.Println("-----------------------------")
	// 7.2.3 メソッドは関数
	{
		a := Adder{start: 10}
		fmt.Println(a.AddTo(5)) // 15
		f1 := a.AddTo           // Adder型の変数aのAddToメソッドをf1に代入
		fmt.Println(f1(10))     // 20
		// メソッド式 -> 第一引数がメソッドのレシーバになる
		f2 := Adder.AddTo
		fmt.Println(f2(a, 15)) // 25 レシーバとしてaを渡す
	}
	// 7.2.7 iotaと列挙型
	{
		
	}
}
