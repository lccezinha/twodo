package usecases

import (
	"reflect"
	"testing"

	"github.com/lccezinha/twodo/internal/twodo"
	stubs "github.com/lccezinha/twodo/internal/twodo/test"
)

func TestCreateService(t *testing.T) {
	t.Run("Given blank title, it returns invalid fields", func(t *testing.T) {
		repository := stubs.NewFakeRepository()
		usecase := NewCreateTodoUseCase(repository)
		presenter := stubs.NewFakePresenter()

		usecase.Run("", "Description", presenter)

		expectedErrs := []twodo.ValidationError{
			twodo.ValidationError{
				Field:   "Title",
				Message: "Can not be blank",
				Type:    "Required",
			},
		}

		if !reflect.DeepEqual(presenter.Errs, expectedErrs) {
			t.Errorf("Expected %v to be eq to %v", presenter.Errs, expectedErrs)
		}
	})

	t.Run("Given valid args, create and return todo", func(t *testing.T) {
		repository := stubs.NewFakeRepository()
		usecase := NewCreateTodoUseCase(repository)
		presenter := stubs.NewFakePresenter()
		todo := twodo.Todo{
			Title:       "Title",
			Description: "Description",
		}

		usecase.Run(todo.Title, todo.Description, presenter)

		if repository.Todo != todo {
			t.Errorf("Expected %v to be eq to %v", repository.Todo, todo)
		}

		if presenter.Todo != todo {
			t.Errorf("Expected %v to be eq to %v", presenter.Todo, todo)
		}
	})
}
