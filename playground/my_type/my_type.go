package my_type

import "fmt"

func My_type() {
	// ゼロ値(いわゆるデフォルト値)が存在する

	/* リテラル */
	/* 整数型 */
	// 整数型のゼロ値は0
	fmt.Printf("1: 整数リテラル %d\n", 10)

	// int型はCPUニヨッテ42bitが64bitに変わる

	//　演算時の型変換は上手く行われない
	fmt.Printf("2: 整数型と浮動小数点数の足し算 %v\n", 1+1.0)
	//=> 22

	// 整数同士の割り算は整数に丸められる（切り捨て）
	fmt.Printf("3: 整数同士の割り算 10 / 3 = %d\n", 10/3)
	//=> 3

	/* 浮動小数点数 */
	fmt.Printf("4: 浮動小数点数リテラル %f\n", 123.456)

	/* rune */
	// コードポイントを表す
	// シングルクオートで囲う。いわゆる文字列とは扱いが違う
	fmt.Printf("5: runeリテラル %U\n", 'a')
	fmt.Printf("5: runeリテラル %U\n", '\n')

	/* 文字列　*/
	// 文字列はバイトの列で構成されている
	// ダブルクオートで囲うと解釈済み文字列リテラル
	fmt.Printf("6: 文字列リテラル %s\n", "Hello World")
	// バッククオートで囲うとロー文字列リテラル。改行出来たりする
	fmt.Printf("6: 文字列リテラル %s\n", `I'm a
	loser.`)

	var str string = "Hello Wolrd"
	var b byte = str[6]

	fmt.Printf("7: str[6]の値 %v\n", b)

	/* 真偽値 */
	fmt.Printf("8: 真偽値 %v\n", true)

	fmt.Printf("9: こんにちは のバイト数 %v\n", len("こんにちは"))
	//=> 9: こんにちは のバイト数 15

	/* 型キャスト */
}
