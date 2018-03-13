package todo

import (
	"reflect"
	"testing"

	"github.com/lccezinha/twodo/internal/repository/todo"
	"github.com/lccezinha/twodo/internal/twodo"
)

func TestCreateService(t *testing.T) {
	t.Run("Given blank title, it returns invalid fields", func(t *testing.T) {
		repository := todo.NewMemoryRepository()
		service := NewCreateService(repository)
		err := service.Run("", "Description")
		expectedErrs := []twodo.ValidationError{
			twodo.ValidationError{
				Field:   "Title",
				Message: "Can not be blank",
				Type:    "Required",
			},
		}

		if !reflect.DeepEqual(err, expectedErrs) {
			t.Errorf("Expected %v to be eq to %v", err, expectedErrs)
		}
	})
}
