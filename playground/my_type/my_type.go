package my_type

import "fmt"

func My_type() {
	// ゼロ値(いわゆるデフォルト値)が存在する

	/* リテラル */
	/* 整数型 */
	// 整数型のゼロ値は0
	fmt.Printf("1: 整数リテラル %d\n", 10)

	//　演算時の型変換は上手く行われない
	fmt.Printf("2: 整数型と浮動小数点数の足し算 %v\n", 1+1.0)
	//=> 22

	/* 浮動小数点数 */
	fmt.Printf("2: 浮動小数点数リテラル %f\n", 123.456)

	/* rune */
	// 文字を表す
	// シングルクオートで囲う。いわゆる文字列とは扱いが違う
	fmt.Printf("3: runeリテラル %U\n", 'a')
	fmt.Printf("3: runeリテラル %U\n", '\n')

	/* 文字列　*/
	// ダブルクオートで囲うと解釈済み文字列リテラル
	fmt.Printf("4: 文字列リテラル %s\n", "Hello World")
	// バッククオートで囲うとロー文字列リテラル。改行出来たりする
	fmt.Printf("4: 文字列リテラル %s\n", `I'm a
	loser.`)

	/* 真偽値 */
	fmt.Printf("5: 真偽値 %v\n", true)
}
