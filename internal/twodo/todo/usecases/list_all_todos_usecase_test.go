package usecases

import (
	"reflect"
	"testing"

	"github.com/lccezinha/twodo/internal/test/fakes"
	"github.com/lccezinha/twodo/internal/twodo"
)

func TestListAllUseCase(t *testing.T) {
	t.Run("When there is no todo to list", func(t *testing.T) {
		repository := fakes.NewFakeRepository()
		usecase := NewListAllUseCase(repository)
		presenter := fakes.NewFakePresenter()

		usecase.Run(presenter)

		if !presenter.ListArgCalled {
			t.Error("Did not call presenter")
		}

		if len(presenter.List) > 1 {
			t.Error("It must not render any todo")
		}
	})

	t.Run("Listing todos from most recente to most old", func(t *testing.T) {
		repository := fakes.NewFakeRepository()
		usecase := NewListAllUseCase(repository)
		presenter := fakes.NewFakePresenter()
		todos := []twodo.Todo{
			twodo.Todo{ID: 1, Description: "Todo #1"},
			twodo.Todo{ID: 2, Description: "Todo #2"},
		}

		repository.ListAllResult = todos
		usecase.Run(presenter)

		if !presenter.ListArgCalled {
			t.Error("Did not call presenter")
		}

		if len(presenter.List) == 0 {
			t.Error("Must render at least one todo")
		}

		if !reflect.DeepEqual(presenter.List, todos) {
			t.Errorf("Expected: %v, received: %v", presenter.List, todos)
		}
	})
}
