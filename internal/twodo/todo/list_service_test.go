package todo

import (
	"testing"

	"github.com/lccezinha/twodo/internal/repository/todo"
	"github.com/lccezinha/twodo/internal/twodo"
)

func TestListService(t *testing.T) {
	t.Run("when there aren't todos return empty", func(t *testing.T) {
		repository := todo.NewMemoryRepository()
		service := NewListService(repository)

		todos, err := service.Run()

		if len(todos) > 0 {
			t.Errorf("Must return 0 todos")
		}

		if err != nil {
			t.Errorf("Error %s", err.Error())
		}
	})

	t.Run("when there are todos return all", func(t *testing.T) {
		repository := todo.NewMemoryRepository()
		service := NewListService(repository)
		serviceCreate := NewCreateService(repository)

		todo := &twodo.Todo{
			Title:       "Title #1",
			Description: "Description #1",
		}

		otherTodo := &twodo.Todo{
			Title:       "Title #2",
			Description: "Description #2",
		}

		serviceCreate.Run(todo.Title, todo.Description)
		serviceCreate.Run(otherTodo.Title, otherTodo.Description)

		todos, err := service.Run()

		if len(todos) == 0 {
			t.Errorf("Must return at least one todo")
		}

		if todos[0].Title != todo.Title {
			t.Errorf("Does not fetch correct todos")
		}

		if todos[1].Title != otherTodo.Title {
			t.Errorf("Does not fetch correct todos")
		}

		if err != nil {
			t.Errorf("Error %s", err.Error())
		}
	})
}
