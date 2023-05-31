package my_pointer

import "fmt"

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func update100(px *int) {
	// デリファレンスで100に書き換える
	*px = 100
}

// 同じ参照先を持ったポインターがコピーされている
func show_address(s *string) {
	fmt.Printf("show_address: sの値 = %v\n", s)
	//=> show_address: sの値 = 0xc00005a250
	fmt.Printf("show_address: sのアドレス = %v\n", &s)
	//=> show_address: sのアドレス = 0xc000006060
}

func change_address(s *string) {
	temp := "hello"
	s = &temp
}

func change_address_2(s *string) {
	temp := "hello"
	// デリファレンスで参照先を書き換える
	*s = temp
}

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

	// nilポインター
	var f *int

	failedUpdate(f)

	// nilポインターは更新できない
	fmt.Printf("2: f = %v\n", f)
	//=> 2: f = <nil>

	x2 := 10

	update100(&x2)

	fmt.Printf("3: update100に渡した後のx2 = %v\n", x2)

	// マップは構造体へのポインターとして実装されている
	// → 関数にマップを渡すということは、ポインターをコピーすることと同じ

	var x3 = 100
	fmt.Printf("4: x3のアドレス = %v\n", &x3)
	//=> 4: x3のアドレス = 0xc00001a0d8

	// ポインター変数として定義
	var px3 *int
	fmt.Printf("4: px3の値 = %v\n", px3)
	//=> 4: px3の値 = <nil>
	fmt.Printf("4: px3のアドレス = %v\n", &px3)
	//=> 4: px3のアドレス = 0xc000006040

	px3 = &x3
	fmt.Printf("5: px3の値 = %v\n", px3)
	//=> 5: px3の値 = 0xc00001a0d8
	fmt.Printf("5: px3のデリファレンス = %v\n", *px3)
	//=> 5: px3のデリファレンス = 100

	// ダブルポインター

	var x4 = 100
	fmt.Printf("6: x4のアドレス = %v\n", &x4)
	//=> 6: x4のアドレス = 0xc00001a0e0

	px4 := &x4
	fmt.Printf("6: pxの値 = %v\n", px4)
	//=> 6: pxの値 = 0xc00001a0e0
	fmt.Printf("6: px4のアドレス = %v\n", &px4)
	//=> 6: px4のアドレス = 0xc000006048

	ppx4 := &px4
	fmt.Printf("6: ppx4の値 = %v\n", ppx4)
	//=> 6: ppx4の値 = 0xc000006048
	fmt.Printf("6: ppx4のアドレス = %v\n", &ppx4)
	//=> 6: ppx4のアドレス = 0xc000006050

	fmt.Printf("6: ppx4のデリファレンス = %v\n", *ppx4)
	//=> 6: ppx4のデリファレンス = 0xc00001a0e0
	fmt.Printf("6: ppx4の2つデリファレンス = %v\n", **ppx4)
	//=> 6: ppx4の2つデリファレンス = 100

	// 値渡し
	// Goには値渡ししかない

	// 参照私っぽいのは参照の値渡し

	name := "kento"
	fmt.Printf("7: nameのアドレス = %v\n", &name)
	//=> 7: nameのアドレス = 0xc00005a250

	pname := &name
	fmt.Printf("7: pnameのアドレス = %v\n", &pname)
	//=> 7: pnameのアドレス = 0xc000006058

	show_address(pname)

	change_address(pname)
	fmt.Printf("8: nameの値 = %v\n", *pname)
	//=> 8: nameの値 = kento

	change_address_2(pname)
	fmt.Printf("8: nameの値 = %v\n", *pname)
	//=> 8: nameの値 = hello
}
