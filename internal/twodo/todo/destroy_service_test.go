package todo

import (
	"testing"

	"github.com/lccezinha/twodo/internal/repository/todo"
	"github.com/lccezinha/twodo/internal/twodo"
)

func TestDestroyService(t *testing.T) {
	t.Run("when todo ID does not exist raise error", func(t *testing.T) {
		repository := todo.NewMemoryRepository()
		service := NewDestroyService(repository)

		err := service.Run(1)

		if err == nil {
			t.Errorf("Must raise error")
		}
	})

	t.Run("when destroy a todo", func(t *testing.T) {
		repository := todo.NewMemoryRepository()
		service := NewDestroyService(repository)
		listService := NewListService(repository)
		createService := NewCreateService(repository)

		todo := &twodo.Todo{
			Title:       "Title #1",
			Description: "Description #1",
		}

		otherTodo := &twodo.Todo{
			Title:       "Title #2",
			Description: "Description #2",
		}

		createService.Run(todo.Title, todo.Description)
		createService.Run(otherTodo.Title, otherTodo.Description)

		todos, _ := listService.Run()

		if len(todos) != 2 {
			t.Errorf("Some error happen when try to add todos")
		}

		service.Run(1)

		todos, _ = listService.Run()

		if len(todos) != 1 {
			t.Errorf("Some error happen when try to destroy a single todo")
		}

		if todos[0].Title != otherTodo.Title {
			t.Error("Does not destroy the correct todo")
		}
	})
}
