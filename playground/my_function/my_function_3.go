package my_function

import (
	"fmt"
	"sort"
)

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func My_function_3() {
	// 無名関数
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Printf("1: 無名関数 %v\n", j)
		}(i)
	}

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}

	fmt.Printf("■ 初期データ %v\n", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	fmt.Printf("■ 2番目のフィールドでソート %v\n", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Printf("■ 3番目のフィールドでソート %v\n", people)

	twoMake := makeMult(2)
	threeMake := makeMult(3)

	fmt.Printf("%v\n", twoMake(10))
	fmt.Printf("%v\n", threeMake(10))
}
