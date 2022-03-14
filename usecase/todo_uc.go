package usecase

import (
	"github.com/negiseijin/clean-architectur-swagger/domain/model"
	"github.com/negiseijin/clean-architectur-swagger/domain/repository"
	"github.com/negiseijin/clean-architectur-swagger/gen/restapi/operations/todo"
	"github.com/negiseijin/clean-architectur-swagger/infrastructure/persistence"
)

type TodoUsecase interface {
	FindTodo(params todo.GetTodoParams, item interface{}) error // 検索
	ReadTodo(item interface{}) error                            // 全検索
	CreateTodo(item interface{}) error                          // 新規作成
	UpdateTodo(item interface{}) error                          // 更新
	DeleteTodo(item interface{}) error                          // 削除
}

type todoUsecase struct {
	Repo repository.DBRepository
}

// CreateTodo implements repository.TodoRepository
func (u *todoUsecase) CreateTodo(item interface{}) error {
	err := u.Repo.Create(model.Todo{}, item)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTodo implements repository.TodoRepository
func (u *todoUsecase) DeleteTodo(item interface{}) error {
	err := u.Repo.Delete(model.Todo{}, item)
	if err != nil {
		return err
	}
	return nil
}

// FindTodo implements repository.TodoRepository
func (u *todoUsecase) FindTodo(params todo.GetTodoParams, item interface{}) error {
	query := map[string]interface{}{"id": params.TodoID}
	err := u.Repo.Find(model.Todo{}, query, item)
	if err != nil {
		return err
	}
	return nil
}

// ReadTodo implements repository.TodoRepository
func (u *todoUsecase) ReadTodo(item interface{}) error {
	err := u.Repo.Read(model.Todo{}, item)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTodo implements repository.TodoRepository
func (u *todoUsecase) UpdateTodo(item interface{}) error {
	err := u.Repo.Update(model.Todo{}, item)
	if err != nil {
		return err
	}
	return nil
}

func NewTodoUsecase(repo persistence.DBRepository) repository.TodoRepository {
	todoUsecase := todoUsecase{
		Repo: persistence.NewDBRepository(repo.DB)}
	return &todoUsecase
}
