package my_map

import "fmt"

func My_map() {
	// キーがstring型で、値がint型のマップ
	var m1 map[string]int

	fmt.Printf("1: m1 = %v\n", m1)
	//=> 1: m1 = map[]

	// 以下のように値を追加できない
	// m1["id"] = 3

	// マップリテラルを使って定義する
	m2 := map[string]int{
		"id": 10,
	}

	fmt.Printf("2: m2 = %v\n", m2)
	//=> 2: m2 = map[id:10]

	// 要素を追加できる
	m2["age"] = 30
	fmt.Printf("3: m2 = %v\n", m2)
	//=> 3: m2 = map[age:30 id:10]

	// キーが存在するかどうかが、okに代入される
	// Option型に似てる？
	m3 := map[string]int{
		"id":  1,
		"age": 30,
	}

	v, ok := m3["id"]

	fmt.Printf("4: v = %v, キーが存在するかどうか = %v\n", v, ok)
	//=> 4: v = 1, キーが存在するかどうか = true

	v, ok = m3["age"]

	fmt.Printf("4: v = %v, キーが存在するかどうか = %v\n", v, ok)
	//=> 4: v = 30, キーが存在するかどうか = true

	v, ok = m3["hobby"]

	fmt.Printf("4: v = %v, キーが存在するかどうか = %v\n", v, ok)
	//=> 4: v = 0, キーが存在するかどうか = false

	// deleteで削除する
	delete(m3, "id")

	fmt.Printf("5: m3 = %v\n", m3)
	//=> 5: m3 = map[age:30]
}
