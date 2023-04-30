package slice

import "fmt"

func Slice() {
	/* 配列 */
	// 要素数3、0で埋められる
	var arr1 [3]int

	fmt.Printf("1: arr1 %v\n", arr1)

	// {}で明示的に要素を入れて初期化する
	var arr2 = [3]int{1, 2, 3}

	fmt.Printf("2: arr2 %v\n", arr2)

	// len、cap
	fmt.Printf("3: arr1.len() %v, arr1.cap() %v\n", len(arr1), cap(arr1))

	// %Tでデータ型を出力
	fmt.Printf("4: arr1のデータ型 %T\n", arr1)

	/* スライス */
	var s1 []int
	s2 := []int{}

	fmt.Printf("5: s1のデータ型 %T, s2のデータ型 %T\n", s1, s2)

	// 要素を動的に追加できる
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("6: s1 %v\n", s1)

	// スライスにスライスを追加する
	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...)
	fmt.Printf("7: s1にs3を追加 %v\n", s1)

	// make = スライスに応じてメモリーを確保
	// 要素数0、キャパシティー2のスライス
	s4 := make([]int, 0, 2)
	// キャパシティー2を超えた、4つの要素を追加
	s4 = append(s4, 1, 2, 3, 4)
	fmt.Printf("8: s4 %v\n", s4)

	s5 := make([]int, 4, 6)
	s6 := s5[1:3]
	s6[1] = 10
	// 同じ実体を参照しているため、s6を書き換えるとs5も書き変わる
	fmt.Printf("9: s5 %v, s6 %v\n", s5, s6)

	// s5に要素を追加(s6には追加されない)
	s5 = append(s5, 100)
	fmt.Printf("9: s5 %v, s6 %v\n", s5, s6)

	// s6に要素を追加(s5にも追加される)
	s6 = append(s6, 999)
	fmt.Printf("9: s5 %v, s6 %v\n", s5, s6)

	s7 := []int{100, 100, 100}
	s8 := make([]int, 3)
	// s8にs7をコピー
	copy(s8, s7)

	fmt.Printf("10: s7 %v, s8 %v", s7, s8)

	// https://qiita.com/k-penguin-sato/items/daad9986d6c42bdcde90
}
