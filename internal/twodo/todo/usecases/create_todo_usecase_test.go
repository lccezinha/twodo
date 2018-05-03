package usecases

import (
	"reflect"
	"testing"

	"github.com/lccezinha/twodo/internal/test/fakes"
	"github.com/lccezinha/twodo/internal/twodo"
)

func TestCreateUseCase(t *testing.T) {
	t.Run("When has a blank description return invalid fields", func(t *testing.T) {
		repository := fakes.NewFakeRepository()
		usecase := NewCreateTodoUseCase(repository)
		presenter := fakes.NewFakePresenter()

		usecase.Run("", presenter)

		expectedErrs := []twodo.ValidationError{
			twodo.ValidationError{
				Field:   "Description",
				Message: "Can not be blank",
				Type:    "Required",
			},
		}

		if !reflect.DeepEqual(presenter.Errs, expectedErrs) {
			t.Errorf("Expected %v to be eq to %v", presenter.Errs, expectedErrs)
		}
	})

	t.Run("When has valid args create and return todo", func(t *testing.T) {
		repository := fakes.NewFakeRepository()
		usecase := NewCreateTodoUseCase(repository)
		presenter := fakes.NewFakePresenter()
		todo := twodo.Todo{Description: "Description"}

		usecase.Run(todo.Description, presenter)

		if presenter.Todo != todo {
			t.Errorf("Expected %v to be eq to %v", presenter.Todo, todo)
		}
	})
}
