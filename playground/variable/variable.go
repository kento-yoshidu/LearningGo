package variable

import "fmt"

const secret = "abc"

func Variable() {
	var i int

	fmt.Printf("1: var-intで宣言すると初期値は0になる %d \n", i)

	var i2 int = 2

	fmt.Printf("2: 初期値を代入することもできる %d \n", i2)

	var i3 = 3

	fmt.Printf("3: 初期値を代入すれば、型推論が働く %d\n", i3)

	var i4 = float64(10)

	fmt.Printf("4: 型キャストを行える。％Tで型を出力する %v %T\n", i4, i4)

	fmt.Printf("5: ❌ int * float64など、型が違うとコンパイルエラー i3 * i4")

	fmt.Printf("5: float(i3)という風に型キャストすれば演算できる float64(i3)*i4=%f\n", float64(i3)*i4)

	i = 100

	fmt.Printf("6: 変数はミュータブルであり、値を書き換えることができる i=%d\n", i)
}
