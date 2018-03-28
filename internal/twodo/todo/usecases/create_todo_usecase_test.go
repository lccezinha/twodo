package usecases

import (
	"reflect"
	"testing"

	"github.com/lccezinha/twodo/internal/repository/todo"
	"github.com/lccezinha/twodo/internal/twodo"
)

type FakePresenter struct {
	todo twodo.Todo
	errs []twodo.ValidationError
}

func (fp *FakePresenter) PresentCreatedTodo(todo twodo.Todo) {
	fp.todo = todo
}

func (fp *FakePresenter) PresentErrors(errs []twodo.ValidationError) {
	fp.errs = errs
}

type FakeRepository struct {
	todo twodo.Todo
}

func NewFakePresenter() *FakePresenter {
	return new(FakePresenter)
}

func (fr *FakeRepository) Save(t twodo.Todo) (twodo.Todo, error) {
	fr.todo = t

	return fr.todo, nil
}

func NewFakeRepository() *FakeRepository {
	return new(FakeRepository)
}

func TestCreateService(t *testing.T) {
	t.Run("Given blank title, it returns invalid fields", func(t *testing.T) {
		repository := todo.NewMemoryRepository()
		usecase := NewCreateTodoUseCase(repository)
		presenter := NewFakePresenter()

		usecase.Run("", "Description", presenter)

		expectedErrs := []twodo.ValidationError{
			twodo.ValidationError{
				Field:   "Title",
				Message: "Can not be blank",
				Type:    "Required",
			},
		}

		if !reflect.DeepEqual(presenter.errs, expectedErrs) {
			t.Errorf("Expected %v to be eq to %v", presenter.errs, expectedErrs)
		}
	})

	t.Run("Given valid args, create and return todo", func(t *testing.T) {
		repository := NewFakeRepository()
		usecase := NewCreateTodoUseCase(repository)
		presenter := NewFakePresenter()
		todo := twodo.Todo{
			Title:       "Title",
			Description: "Description",
		}

		usecase.Run(todo.Title, todo.Description, presenter)

		if repository.todo != todo {
			t.Errorf("Expected %v to be eq to %v", repository.todo, todo)
		}

		if presenter.todo != todo {
			t.Errorf("Expected %v to be eq to %v", presenter.todo, todo)
		}
	})
}
