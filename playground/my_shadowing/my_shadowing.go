package my_shadowing

import (
	"fmt"
	"math/rand"
	"time"
)

func My_shadowing() {
	x := 8

	// ifブロック
	if x < 10 {
		fmt.Printf("1: ifブロックの中でxを出力 = %v\n", x)
		x := 5
		fmt.Printf("1: ifブロックの中でxを定義 = %v\n", x)
	}

	fmt.Printf("1: ifブロックの外のx = %v\n", x)

	rand.Seed(time.Now().Unix())

	// nを定義、elseブロックまで利用できる
	if n := rand.Intn(10); n < 5 {
		fmt.Printf("2: nの値(小さい) %v\n", n)
	} else if n > 10 {
		fmt.Printf("2: nの値(大きい) %v\n", n)
	} else {
		fmt.Printf("2: nの値(普通くらい) %v\n", n)
	}
	// ここでは使えない
	// fmt.Printf("🐀❌ nの値 %v", n)

	// Goではループ構文はforだけ

	// 初期化にvarは使えない
	for i := 0; i < 10; i++ {
		fmt.Printf("3: iの値 = %v\n", i)
	}

	// 条件のみ(whileっぽく書ける)
	i := 1

	for i < 10 {
		fmt.Printf("4: i = %v\n", i)
		i = i + 1
	}

	// 無限ループ
	i2 := 0

	for {
		if i2 == 10 {
			break
		}
		i2++
		fmt.Printf("5: i = %v\n", i2)
	}

	for i := 1; i <= 30; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("FizzBuzz [%v]\n", i)
			continue
		}
		if i%3 == 0 {
			fmt.Printf("Fizz [%v]\n", i)
			continue
		}
		if i%5 == 0 {
			fmt.Printf("Buzz [%v]\n", i)
			continue
		}
	}

	// for-range
	// 複合型のみ使える
	// for-ofみたいなやつ

	s1 := []int{2, 4, 6, 8, 10}

	for index, value := range s1 {
		fmt.Printf("6: index is %v, value is %v\n", index, value)
	}

	m1 := map[string]int{"takashi": 30, "keiko": 42, "shinji": 18}

	for key := range m1 {
		fmt.Printf("7: mapのキーのみ出力します %v\n", key)
	}

	// 文字列を回す

	for i, char := range "Hello World" {
		fmt.Printf("8: 文字列の%v字目は%vです。\n", i+1, string(char))
	}

	// 値はコピーされる
	s2 := []int{1, 2, 3, 4, 5}

	for _, v := range s2 {
		v *= 2
		fmt.Printf("9: forループの中で2倍にした数 %v\n", v)
	}

	// 影響を受けない
	fmt.Printf("9: s2 = %v\n", s2)
}
