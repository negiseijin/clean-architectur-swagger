package repository

import "github.com/negiseijin/clean-architectur-swagger/gen/restapi/operations/todo"

/*
Todoリポジトリ
*/
type TodoRepository interface {
	FindTodo(params todo.GetTodoParams, item interface{}) error // 検索
	ReadTodo(item interface{}) error                            // 全検索
	CreateTodo(item interface{}) error                          // 新規作成
	UpdateTodo(item interface{}) error                          // 更新
	DeleteTodo(item interface{}) error                          // 削除
}
