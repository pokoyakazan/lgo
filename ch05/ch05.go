package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

// 5.1 関数の宣言と呼び出し
func div(numerator int, denominnator int) int {
	if denominnator == 0 {
		return 0
	}
	return numerator / denominnator
}

// 5.1.2 可変町引数とスライス
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// 5.1.3 複数の戻り値
func divAndRemainder(numerator, denominnator int) (int, int, error) {
	// または 5.1.6 名前付き戻り値 ではこう書ける
	// func divAndRemainder(numerator, denominnator int) (result int, remainder int, err error) {
	if denominnator == 0 {
		return 0, 0, errors.New("denominator is zero")
	}
	return numerator / denominnator, numerator % denominnator, nil
}

// 5.3.1 関数引数
type Person2 struct {
	FirstName string
	LastName  string
	Age       int
}

// 5.3.2 関数から関数を返す
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

// 5.5 Go は値渡し
type person struct {
	age  int
	name string
}

func modifyFailes(i int, s string, p person) {
	i *= 100
	s = "さようなら"
	p.name = "Bob"
}

func modMap(m map[int]string) {
	m[2] = "こんにちは"
	m[3] = "さようなら"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)

}

func main() {
	// 5.1 関数の宣言と呼び出し
	{
		result := div(10, 2)
		println(result)
		fmt.Println("---")
		// 5.1.2 可変町引数とスライス
		fmt.Println(addTo(3))             // []
		fmt.Println(addTo(3, 2))          // [5]
		fmt.Println(addTo(3, 2, 4, 6, 8)) // [5 7 9 11]
		a := []int{4, 3}
		fmt.Println(addTo(3, a...))           // [7 6]
		fmt.Println(addTo(3, []int{4, 3}...)) // [7 6]
		fmt.Println("---")
		// 5.1.3 複数の戻り値
		result, remainder, err := divAndRemainder(10, 3)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(result, remainder) // 3 1
		fmt.Println("-----------------------------")
		// 5.2.2 無名関数
		for i := 0; i < 5; i++ {
			a := func(j int) int {
				return j * j
			}(i)
			fmt.Println(a)
		}
	}

	fmt.Println("-----------------------------")

	// 5.3 クロージャ
	{
		// 5.3.1 関数引数
		type Person struct {
			FirstName string
			LastName  string
			Age       int
		}
		people := []Person{
			{"Hanako", "Yamada", 20},
			{"Sakura", "Takahashi", 25},
			{"Taro", "Yamamoto", 30},
		}
		fmt.Println(people) // [{Hanako Yamada 20} {Sakura Takahashi 25} {Taro Yamamoto 30}]
		// LastNameでソート
		sort.Slice(people, func(i, j int) bool {
			return people[i].LastName < people[j].LastName
		})
		fmt.Println(people) // [{Sakura Takahashi 25} {Hanako Yamada 20} {Taro Yamamoto 30}]
		fmt.Println("---")

		// 5.3.2 関数から関数を返す
		twoBase := makeMult(2)   // 2倍する関す
		threeBase := makeMult(3) // 3倍する関数
		for i := 0; i <= 5; i++ {
			fmt.Println(i, ": ", twoBase(i), ",", threeBase(i))
		}

		fmt.Println("-----------------------------")

		// 5.4 defer
		{
			file_name := "test.txt"
			f, err := os.Open(file_name)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			data := make([]byte, 2048)
			for { // 無限ループ
				count, err := f.Read(data)    // 読み込んだバイト数とエラーを返す
				os.Stdout.Write(data[:count]) // 標準出力に出力
				if err != nil {
					if err != io.EOF { // ファイルの終わりでないならば
						log.Fatal(err) // エラーを出力して終了
					}
				}
				break
			}

		}
		fmt.Println("-----------------------------")

		// 5.5 Go は値渡し
		{
			p := person{}
			i := 2
			s := "こんにちは"
			fmt.Println(i, s, p) // 2 こんにちは {0 }
			modifyFailes(i, s, p)
			fmt.Println(i, s, p) // 2 こんにちは {0 }

			fmt.Println("---")

			// マップとスライスはポインタで実装されているので参照渡し
			m := map[int]string{
				1: "1番目",
				2: "2番目",
			}
			modMap(m)
			fmt.Println(m) // map[2:こんにちは 3:さようなら]

			ss := []int{1, 2, 3}
			modSlice(ss)
			fmt.Println(ss) // [2 4 6]

			fmt.Println("---")
		}

	}

}
