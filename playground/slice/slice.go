package slice

import "fmt"

func Slice() {
	/* 配列 */

	// 配列の宣言
	// int型で要素数3という意味。ゼロ値(0)で埋められる
	var arr1 [3]int

	fmt.Printf("1: arr1 %v\n", arr1)
	//=> arr1 [0 0 0]

	// 配列リテラル
	// {}で明示的に要素を入れて初期化する
	arr2 := [3]int{1, 2, 3}

	fmt.Printf("2: arr2 = %v\n", arr2)

	// len、cap
	fmt.Printf("3: arr1.len() %v, arr1.cap() %v\n", len(arr1), cap(arr1))

	// %Tでデータ型を出力
	fmt.Printf("4: arr1のデータ型 %T\n", arr1)

	/* スライス */
	// 宣言時に大きさを指定しない
	var s1 []int
	s2 := []int{}

	fmt.Printf("5: s1のデータ型 %T, s2のデータ型 %T\n", s1, s2)
	//=> 5: s1のデータ型 []int, s2のデータ型 []int

	fmt.Printf("6: s1はnilと等しい %v", s1 == nil)

	// appendを使って要素を動的に追加できる
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("6: s1 %v\n", s1)

	// スライスにスライスを追加する
	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...)
	fmt.Printf("7: s1にs3を追加 %v\n", s1)

	/* make */
	// スライスは自動でキャパシティを追加してくれるが、最大サイズが分かっているならmakeを使う方が効率的

	// 要素数0、キャパシティー2のスライス
	s4 := make([]int, 0, 2)
	// 要素数2、キャパシティー2のスライス(0で埋められる)
	s5 := make([]int, 2, 2)

	fmt.Printf("8: s4 = %v\n", s4)
	//=> 8: s4 = []
	fmt.Printf("8: s5 = %v\n", s5)
	//=> 8: s5 = [0 0]

	// s4にappendで追加してみる
	s4 = append(s4, 100)

	fmt.Printf("9: appendした後のs4 = %v\n", s4)
	//=> 9: appendした後のs4 = [100

	// s5は[0 0]なので、appendで追加すると要素数は3になる
	s5 = append(s5, 100)

	fmt.Printf("9: appendした後のs5 = %v\n", s5)
	//=> 9: appendした後のs5 = [0 0 100]

	/* スライス式 */
	// スライスから別のスライスを作る
	// [m:n]の形
	s6 := []int{1, 2, 3, 4}
	s6a := s6[:2] // 最初から2番目まで
	fmt.Printf("10: s6の[:2] %v\n", s6a)
	//=> 10: s6の[:2] [1 2]

	s6b := s6[0:] // index0から最後まで
	fmt.Printf("10: s6の[0:] %v\n", s6b)
	//=> 10: s6の[0:] [1 2 3 4]

	// 存在しない範囲を指定するとエラーになる
	// panic: runtime error: slice bounds out of range [:10] with capacity 4
	// s6c := s6[:10]
	// fmt.Printf("10: s6の[:10], %v\n", s6c)

	// スライスからスライスを切り出す作業はメモリーを共有している
	s7 := []int{1, 2, 3, 4}
	s7a := s7[:]

	fmt.Printf("11: s7 = %v, s7a = %v\n", s7, s7a)
	//=> 11: s7 = [1 2 3 4], s7a = [1 2 3 4]

	// s7aの一部を書き換える
	s7a[0] = 10

	// 再度出力すると、両方の値が変わっていることが分かる
	fmt.Printf("11: s7 = %v, s7a = %v\n", s7, s7a)
	//=> 11: s7 = [10 2 3 4], s7a = [10 2 3 4]

	s8 := []int{1, 2, 3, 4}
	s8a := s8[:]

	// appendで追加する
	s8a = append(s8a, 100)

	fmt.Printf("12: s8 = %v, s8a = %v\n", s8, s8a)
	//=> 12: s8 = [1 2 3 4], s8a = [1 2 3 4 100]

	// スライスのコピー
	s9 := []int{1, 2, 3, 4}
	s9a := make([]int, 4)

	copy(s9a, s9)

	fmt.Printf("13: s9 = %v, s9a = %v\n", s9, s9a)

	s9a[0] = 100

	fmt.Printf("13: s9 = %v, s9a = %v\n", s9, s9a)
	//=> 13: s9 = [1 2 3 4], s9a = [100 2 3 4]
	/*
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

		fmt.Printf("10: s7 %v, s8 %v\n", s7, s8)

		// https://qiita.com/k-penguin-sato/items/daad9986d6c42bdcde90

		/* map */

	// キーの型と値の型を指定する
	var m1 map[string]int
	m2 := map[string]int{}

	// varを使うとnilと等しい
	fmt.Printf("11: m1 = nil ? %v\n", m1 == nil)
	// :=を使うとnilと等しくない
	fmt.Printf("11: m2 = nil ? %v\n", m2 == nil)

	m2["id"] = 1
	m2["age"] = 30

	fmt.Printf("12: m2[\"id\"] = %v, m2[\"age\"] = %v\n", m2["id"], m2["age"])

	// deleteで削除
	delete(m2, "age")

	fmt.Printf("12: age削除後のm2 = %v\n", m2)

	m2["age"] = 30
	m2["height"] = 180

	// for文で回す
	for k, v := range m2 {
		fmt.Printf("13: m2[\"%v\"] = %v\n", k, v)
	}
}
