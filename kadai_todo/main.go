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
func (todo *ToDo) printTodos(todos []ToDo) {
	for _, v := range todos {
		status := "未完了"
		if v.Completed {
			status = "完了"
		}

		fmt.Printf("[%s](ID:%d)%s\n", status, v.Id, v.Title)
	}
}

func main() {
	var todos []ToDo = []ToDo{
		{1, "学習計画", false},
		{2, "環境構築", false},
		{3, "基礎文法", false},
	}

	var t ToDo
	t.printTodos(todos)

	todos[0].Complete()
	todos[1].Complete()

	t.printTodos(todos)

}