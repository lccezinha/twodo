package usecases

import (
	"testing"

	"github.com/lccezinha/twodo/internal/test/fakes"
	"github.com/lccezinha/twodo/internal/twodo"
)

func TestDestroyUseCase(t *testing.T) {
	t.Run("Given an ID of a todo, destroy it", func(t *testing.T) {
		repository := fakes.NewFakeRepository()
		usecase := NewDestroyUseCase(repository)
		presenter := fakes.NewFakePresenter()
		todo := twodo.Todo{ID: 1, Description: "Description"}

		usecase.Run(todo.ID, presenter)

		if !presenter.DestroyArgCalled {
			t.Error("Did not call presenter")
		}
	})
}
