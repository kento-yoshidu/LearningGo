package variable

import "fmt"

const secret = "abc"

func Variable() {
	/* varを使った変数宣言 */
	var i int = 10
	fmt.Printf("1: i = %d\n", i)

	// 型推論が働くのでintは省略できる
	var i2 = 10
	fmt.Printf("2: i2 = %d\n", i2)

	// 変数名と型だけ指定して、後から代入することもできる
	// この時、ゼロ値が代入される
	var i3 int

	fmt.Printf("3: ゼロ値が代入される %d \n", i3)

	// 関数内に限り、:=を使用することができる

	i4 := 10
	fmt.Printf("4: := で代入した値 %d\n", i4)

	/*
		Unlike regular variable declarations, a short variable declaration may redeclare variables provided they were originally declared earlier in the same block (or the parameter lists if the block is the function body) with the same type, and at least one of the non-blank variables is new.
		As a consequence, redeclaration can only appear in a multi-variable short declaration.
		Redeclaration does not introduce a new variable; it just assigns a new value to the original.
	*/

	// 複数代入の時の左辺に新しい変数がある場合、宣言済み変数に再代入できる
	// なんでこんな仕様になっているのか分からない
	i4, x := 100, 200

	fmt.Printf("5: 再代入したi4の値とxの値 %d, %d", i4, x)

	fmt.Printf("4: 型キャストを行える。％Tで型を出力する %v %T\n", i4, i4)

	// fmt.Printf("5: ❌ int * float64など、型が違うとコンパイルエラー i3 * i4")

	// fmt.Printf("5: float(i3)という風に型キャストすれば演算できる float64(i3)*i4=%f\n", float64(i4)*i5)

	i = 100

	fmt.Printf("6: 変数はミュータブルであり、値を書き換えることができる i=%d\n", i)
}
