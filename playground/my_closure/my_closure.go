package my_closure

import "fmt"

func addGen() func() int {
	x := 0

	return func() int {
		x++

		return x
	}
}

func closureKvs() func(cmd string, key string, value interface{}) interface{} {
	/* 一見、変数storeはローカル変数に見えますが、内部的にはクロージャに属する変数として機能しています。
	** クロージャに属する変数として機能するとは、変数storeが何かしらの形で参照され続ける限り、格納されている値が破棄されることがありません。
	 */
	store := make(map[string]interface{})

	return func(cmd string, key string, value interface{}) interface{} {
		switch cmd {
		case "ADD":
			store[key] = value
			return "OK"
		case "GET":
			getVal := store[key]
			return getVal
		}

		return nil
	}

}

func MyClosure() {
	// addというレキシカルスコープに閉じ込めている
	add := addGen()
	fmt.Printf("addの値 = %v\n", add())
	fmt.Printf("addの値 = %v\n", add())
	fmt.Printf("addの値 = %v\n", add())

	// add2というレキシカルスコープに閉じ込めている
	add2 := addGen()
	fmt.Printf("add2の値 = %v\n", add2())
	fmt.Printf("add2の値 = %v\n", add2())
	fmt.Printf("add2の値 = %v\n", add2())

	// kvsというレキシカルスコープ
	kvs := closureKvs()

	// kvsに値を格納
	fmt.Printf("kvs is %v\n", kvs("ADD", "mykey", "my value"))

	fmt.Printf("値取り出し %v\n", kvs("GET", "mykey", ""))
}
