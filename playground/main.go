package main

import (
	"fmt"
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
}
