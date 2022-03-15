package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/negiseijin/clean-architectur-swagger/gen/models"
	"github.com/negiseijin/clean-architectur-swagger/gen/restapi/operations/todo"
	"github.com/negiseijin/clean-architectur-swagger/infrastructure/persistence"
	"github.com/negiseijin/clean-architectur-swagger/usecase"
)

type TodoHandler interface {
	GetTodoListHandler(params todo.GetTodoListParams) middleware.Responder    // TODOリスト取得
	GetTodoHandler(params todo.GetTodoParams) middleware.Responder            // TODO取得
	PostTodoHandler(params todo.CreateTodoParams) middleware.Responder        // TODO作成
	PutTodoHandler(params todo.UpdateTodoParams) middleware.Responder         // TODO更新
	PutTodoListHandler(params todo.UpdateTodoListParams) middleware.Responder // TODOリスト更新
}

type todoHandler struct {
	Usecase usecase.TodoUsecase
}

// PostTodoHandler implements TodoHandler
func (th *todoHandler) PostTodoHandler(params todo.CreateTodoParams) middleware.Responder {
	th.Usecase.CreateTodo(params.Body)
	return todo.NewCreateTodoCreated()
}

// PutTodoHandler implements TodoHandler
func (th *todoHandler) PutTodoHandler(params todo.UpdateTodoParams) middleware.Responder {
	panic("unimplemented")
}

// PutTodoListHandler implements TodoHandler
func (th *todoHandler) PutTodoListHandler(params todo.UpdateTodoListParams) middleware.Responder {
	panic("unimplemented")
}

// GetTodoHandler implements TodoHandler
func (th *todoHandler) GetTodoHandler(params todo.GetTodoParams) middleware.Responder {
	res := models.Todo{}
	th.Usecase.FindTodo(params, res)
	return todo.NewGetTodoOK().WithPayload(&res)
}

// GetTodoListHandler implements TodoHandler
func (th *todoHandler) GetTodoListHandler(params todo.GetTodoListParams) middleware.Responder {
	res := models.TodoList{}
	th.Usecase.ReadTodo(res)
	return todo.NewGetTodoListOK().WithPayload(res)
}

func NewTodoHandler(repo persistence.DBRepository) TodoHandler {
	todoHandler := todoHandler{
		Usecase: usecase.NewTodoUsecase(repo),
	}
	return &todoHandler
}
