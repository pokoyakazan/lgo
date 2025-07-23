package main

import "fmt"

// 6.3 ポインタはミュータブル（変更可能）の印
func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

// 6.4 ポインタは最後の手段
/*
// 悪い例
func MakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
}
// 良い例
func MakeFoo() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}
*/

func main() {
	// 6.1 ポインタ入門
	{
		x := 10
		pointerToX := &x
		fmt.Println(pointerToX)  // アドレスが表示される
		fmt.Println(*pointerToX) // 10
		z := 5 + *pointerToX
		fmt.Println(z) // 15
		var y *int
		fmt.Println(y) // nil
		// fmt.Println(*y) // panic

		var pointerToX2 *int
		pointerToX2 = &x
		fmt.Println(pointerToX2) // アドレスが表示される
	}
	fmt.Println("-----------------------------")
	// 6.3 ポインタはミュータブル（変更可能）の印
	{
		x := 10
		failedUpdate(&x)
		fmt.Println(x) // 10
		update(&x)
		fmt.Println(x) // 20
	}
	fmt.Println("-----------------------------")
	// 6.4 ポインタは最後の手段
	{

	}
	fmt.Println("-----------------------------")

}
