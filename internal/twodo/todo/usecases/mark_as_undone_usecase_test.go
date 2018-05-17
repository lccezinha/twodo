package usecases

import (
	"testing"

	"github.com/lccezinha/twodo/internal/test/fakes"
	"github.com/lccezinha/twodo/internal/twodo"
)

func TestMarkTodoAsUndone(t *testing.T) {
	repository := fakes.NewFakeRepository()
	usecase := NewMarkAsUndoneUseCase(repository)
	presenter := fakes.NewFakePresenter()
	todo := twodo.Todo{ID: 1, Description: "Description", Done: true}

	usecase.Run(todo.ID, presenter)

	if !presenter.UpdateArgCalled {
		t.Error("Expect to touch UpdatedArgCalled")
	}
}
