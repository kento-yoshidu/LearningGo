package my_function

import (
	"fmt"
	"strconv"
)

func add(i, j int) int {
	return i + j
}

func sub(i, j int) int {
	return i - j
}

func mul(i, j int) int {
	return i * j
}

func div2(i, j int) int {
	return i / j
}

func My_function_2() {
	var opMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div2,
	}

	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "*", "three"},
	}

	for _, expression := range expressions {
		// 最初の引数を数値に変換
		p1, err := strconv.Atoi(expression[0])

		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}

		op := expression[1]

		// 関数を変数に代入
		opFunc, ok := opMap[op]

		if !ok {
			fmt.Print(expression, " -- ", "定義されていない演算子です : ", op, "\n")
			continue
		}

		p2, err := strconv.Atoi(expression[2])

		if err != nil {
			continue
		}

		result := opFunc(p1, p2)

		fmt.Printf("1: %v %v %v = %v\n", p1, op, p2, result)
	}
}
