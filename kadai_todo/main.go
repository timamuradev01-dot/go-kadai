package main

import "fmt"

// 
type ToDo struct {
	Id  int // 各ToDoアイテムを識別するためのID
	Title string    // ToDoアイテムのタイトル
	Completed bool    // ToDoアイテムが完了したかどうか（true=完了）
}

// 
func (todo *ToDo) Complete() {
	fmt.Printf("ID：%dのToDoを完了に更新します\n", todo.Id)
	todo.Completed = true
}


// 
func printTodos(todos []ToDo) {
	for _, v := range todos {
		status := "未完了"
		if v.Completed {
			status = "完了"
		}

		fmt.Printf("[%s](ID:%d)%s\n", status, v.Id, v.Title)
	}
}

func main() {
	// var todos []ToDo = []ToDo{
	// 	{1, "学習計画", false},
	// 	{2, "環境構築", false},
	// 	{3, "基礎文法", false},
	// }

	// var t ToDo

	// 1. 空のスライスを作成
	todos := []ToDo{}

	// 2. 初期値で ToDo を3つ作成
	todo1 := ToDo{Id: 1, Title: "学習計画", Completed: false}
	todo2 := ToDo{Id: 2, Title: "環境構築", Completed: false}
	todo3 := ToDo{Id: 3, Title: "基礎文法", Completed: false}

	// 3. スライスに追加
	todos = append(todos, todo1, todo2, todo3)
	printTodos(todos)

	todos[0].Complete()
	todos[1].Complete()

	printTodos(todos)

}