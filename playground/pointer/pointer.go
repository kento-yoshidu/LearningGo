package pointer

import (
	"fmt"
	"unsafe"
)

func Pointer() {
	// 0が代入される
	var ui1 uint16

	fmt.Printf("1: ui1のメモリーアドレス %p\n", &ui1)

	var ui2 uint16
	fmt.Printf("1: ui2のメモリーアドレス %p\n", &ui2)

	// ポインター変数
	// ポインターはある変数の先頭アドレスを持つ = 8byte
	var p1 *uint16

	fmt.Printf("2: p1の値 %v\n", p1)

	// ui1の先頭アドレスを代入
	p1 = &ui1
	fmt.Printf("2: 代入した後のp1の値 %v\n", p1)
	fmt.Printf("2: p1の先頭アドレス %p\n", &p1)
	fmt.Printf("2: p1のメモリーサイズ %dbytes\n", unsafe.Sizeof(p1))
	fmt.Printf("2: デリファレンス(p1が指すui1の値) %v\n", *p1)

	// ui1を書き換え
	*p1 = 1
	fmt.Printf("3: 書き換えた後のui1の値 %v\n", ui1)

	// ダブルポインター
	var pp1 **uint16 = &p1

	fmt.Printf("4: pp1の値 %v\n", pp1)
	fmt.Printf("4: p1の先頭アドレス %p\n", &p1)
	fmt.Printf("4: pp1の先頭アドレス %p\n", &pp1)
	fmt.Printf("4: pp1のメモリーサイズ %dbyte\n", unsafe.Sizeof(pp1))

	fmt.Printf("5: pp1のデリファレンス %v\n", *pp1)
	fmt.Printf("5: pp1の2重デリファレンス %v\n", **pp1)

	**pp1 = 2

	fmt.Printf("6: 書き換えた後のui1の値 %v\n", ui1)
}
