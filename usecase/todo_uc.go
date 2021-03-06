package usecase

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/negiseijin/clean-architectur-swagger/domain/model"
	"github.com/negiseijin/clean-architectur-swagger/domain/repository"
	"github.com/negiseijin/clean-architectur-swagger/gen/models"
	"github.com/negiseijin/clean-architectur-swagger/gen/restapi/operations/todo"
	"github.com/negiseijin/clean-architectur-swagger/helpers"
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
	t := item.(*models.CreateTodo)
	value := &model.Todo{
		Base: model.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			// DeletedAt: ,
		},
		Name: t.Name,
		Done: false,
	}

	err := u.Repo.Create(model.Todo{}, value)
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
	results, err := u.Repo.Read(model.Todo{})
	if err != nil {
		return err
	}

	res := []model.Todo{}
	err = helpers.NewHelpers().JsonToStruct(results, &res)
	if err != nil {
		return err
	}

	switch t := item.(type) {
	case *models.TodoList:
		s := models.TodoList{}
		for _, v := range res {
			s = append(s, &models.Todo{
				CreatedAt: strfmt.Date(v.CreatedAt),
				Done:      &v.Done,
				ID:        strfmt.UUID(v.ID.String()),
				Name:      v.Name,
				UpdatedAt: strfmt.Date(v.UpdatedAt),
			})
		}
		*t = s
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
		Repo: persistence.NewDBRepository(repo.DB),
	}
	return &todoUsecase
}
