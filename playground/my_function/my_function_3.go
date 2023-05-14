package my_function

import "fmt"

func My_function_3() {
	// 無名関数
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Printf("1: 無名関数 %v\n", j)
		}(i)
	}

}
