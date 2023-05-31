package my_type

import "fmt"

// 型Person
// 構造体リテラルを基底型としてもつ
type Person struct {
	LastName  string
	FirstName string
	Age       int
}

// Echoメソッド
func (p Person) Echo() string {
	return fmt.Sprintf("%s %s :年齢%d歳", p.LastName, p.FirstName, p.Age)
}

func My_type_2() {
	/* 抽象型と具象型 */

	// 全ての型は、ベースになる型である基底型を持っている
	// → 基本型か型リテラルなら自分自身が基底型
	// → それ以外は、宣言で参照している型が基底型
}
