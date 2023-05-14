package my_function

import (
	"errors"
	"fmt"
	"os"
)

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

// 引数の型が同じときは、型をまとめて書ける
// func dummy(a, b, c int) {}

// 可変長引数
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}

	return out
}

// 複数の戻り値
func divAndRemainder(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("0で割ることはできません。")
	}

	return numerator / denominator, numerator % denominator, nil
}

// 名前付き戻り値
func divAndRemainder2(numerator, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		return numerator, denominator, errors.New("0で割ることはできません。")
	}

	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func My_function() {
	fmt.Printf("1: 10 / 5 = %v\n", div(10, 5))

	fmt.Printf("2: 可変長引数 (3, 2) %v\n", addTo(3, 2))
	//=> 2: 可変長引数 (3, 2) [5]

	fmt.Printf("2: 可変長引数 (3, 2, 5) %v\n", addTo(3, 2, 5))
	//=> 2: 可変長引数 (3, 2, 5) [5 8]

	result, remainder, err := divAndRemainder(5, 2)
	// result, remainder, err := divAndRemainder(10, 0)
	//=> 0で割ることはできません。
	//=> exit status 1

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("3: 商 = %v, 余り = %v\n", result, remainder)

	// ブランク識別子(戻り値を無視できる)
	_, remainder2, _ := divAndRemainder(10, 3)
	fmt.Printf("3: 余り = %v\n", remainder2)

	x, y, z := divAndRemainder2(10, 2)

	if z != nil {
		fmt.Println(z)
		os.Exit(1)
	}

	fmt.Printf("4: x = %v, y = %v, z = %v\n", x, y, z)
}
