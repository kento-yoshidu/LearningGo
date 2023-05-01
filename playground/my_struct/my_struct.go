package my_struct

import "fmt"

type Task struct {
	Title    string
	Estimate int
}

// メソッドの追加
// taskはレシーバーと呼ばれる
func (task Task) extendEstimate() {
	// 受け取ったインスタンスのEstimateに10を加える
	task.Estimate += 10
}

func Struct() {
	// インスタンス生成
	task1 := Task{
		Title:    "Learning Golang",
		Estimate: 3,
	}

	fmt.Printf("1: task1 %v\n", task1)

	// フィールドを書き換える
	task1.Title = "Learning TypeScript"

	fmt.Printf("2: フィールドを書き換えた後のtask1 %v\n", task1)

	var task2 Task = task1

	task2.Title = "Learning Rust"

	fmt.Printf("3: task1.Title = %v, task2.Title = %v\n", task1.Title, task2.Title)

	// ポインター変数
	task1p := &Task{
		Title:    "Programming Lisp",
		Estimate: 10,
	}

	fmt.Printf("4: task1pのデータ型 %v\n", task1p)
	//=> 4: task1pのデータ型 &{Programming Lisp 10}

	fmt.Printf("4: task1pが指しているデータ %v\n", *task1p)
	//=> 4: task1pが指しているデータ {Programming Lisp 10}

	// デリファレンスで値を書き換える
	(*task1p).Title = "Changed"

	fmt.Printf("5: 書き換えた後のtask1p %v\n", task1p)

	var task2p *Task = task1p

	task2p.Title = "Changed ny task2p"

	fmt.Printf("6: task1 %v, task2 %v\n", task1p, task2p)

	task3 := Task{
		Title:    "Task3",
		Estimate: 20,
	}

	task1.extendEstimate()
	task3.extendEstimate()

	fmt.Printf("7: task3 %v\n", task3)
	fmt.Printf("7: task1 %v\n", task1)
}
