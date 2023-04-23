package main

import (
	"fmt"
	"playground/pointer"
	"playground/variable"
)

func main() {
	fmt.Println("--- 変数 ---")

	variable.Variable()

	fmt.Println("--- ポインターとメモリーアドレス ---")

	pointer.Pointer()
}
