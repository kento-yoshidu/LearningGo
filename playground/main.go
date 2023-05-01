package main

import (
	"fmt"
	"playground/my_struct"
	"playground/pointer"
	"playground/slice"
	"playground/variable"
)

func main() {
	fmt.Println("--- 変数 ---")

	variable.Variable()

	fmt.Println("--- ポインターとメモリーアドレス ---")

	pointer.Pointer()

	fmt.Println("--- スライス ---")

	slice.Slice()

	fmt.Println("--- 構造体 ---")

	my_struct.Struct()
}
