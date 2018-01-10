package todo

import (
	"testing"

	"github.com/lccezinha/twodo/internal/repository/todo"
)

func TestCreateService(t *testing.T) {
	t.Run("Fail when title is blank", func(t *testing.T) {
		repository := todo.NewMemoryRepository()
		service := NewCreateService(repository)

		title := ""
		description := "Some description"

		err := service.Run(title, description)

		if err != errCannotBeBlank {
			t.Errorf("Error: %s", err.Error())
		}
	})

	t.Run("Fail description is blank", func(t *testing.T) {
		repository := todo.NewMemoryRepository()
		service := NewCreateService(repository)

		title := "Some title"
		description := ""

		err := service.Run(title, description)

		if err != errCannotBeBlank {
			t.Errorf("Error: %s", err.Error())
		}
	})

	t.Run("Success", func(t *testing.T) {
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
