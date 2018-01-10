package todo

import (
	"testing"

	"github.com/lccezinha/twodo/internal/repository/todo"
)

func TestCreateService(t *testing.T) {
	t.Run("when title is blank must raise error", func(t *testing.T) {
		repository := todo.NewMemoryRepository()
		service := NewCreateService(repository)

		title := ""
		description := "Some description"

		err := service.Run(title, description)

		if err != errCannotBeBlank {
			t.Errorf("Error: %s", err.Error())
		}
	})

	t.Run("when description is blank must raise error", func(t *testing.T) {
		repository := todo.NewMemoryRepository()
		service := NewCreateService(repository)

		title := "Some title"
		description := ""

		err := service.Run(title, description)

		if err != errCannotBeBlank {
			t.Errorf("Error: %s", err.Error())
		}
	})

	t.Run("when all fields are filled must save", func(t *testing.T) {
		repository := todo.NewMemoryRepository()
		service := NewCreateService(repository)

		title := "Some title"
		description := "Some Description"

		err := service.Run(title, description)

		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})
}
