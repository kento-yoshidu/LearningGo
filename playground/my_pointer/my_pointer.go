package my_pointer

import "fmt"

func MyPointer() {
	// メモリーを4バイト確保
	var x = 10

	// 1バイト確保
	var y = true

	// &はアドレス演算子という
	// 変数の前につけることで、その変数のアドレスを返す
	// それぞれの変数のアドレスを記憶している
	// 共にポインター型の変数となっている
	pointerX := &x
	pointerY := &y

	// *は関節参照を行う
	// ポインター型の変数の前につけることで、そのポインターが参照するアドレスの値を返す

	fmt.Printf("1: pointerX = %v, pointerY = %v\n", pointerX, pointerY)
	//=> 1: pointerX = 0xc00001a0a8, pointerY = 0xc00001a0d0
	fmt.Printf("1: それぞれが指すアドレスに格納されている値 x = %v, y = %v\n", *pointerX, *pointerY)
	//=> 1: それぞれが指すアドレスに格納されている値 x = 10, y = true
	fmt.Printf("1: &を付けると、ポインター自身のアドレスを出力する x = %v, y = %v\n", &pointerX, &pointerY)
	//=> 1: &を付けると、ポインター自身のアドレスを出力する x = 0xc000006030, y = 0xc000006038
}