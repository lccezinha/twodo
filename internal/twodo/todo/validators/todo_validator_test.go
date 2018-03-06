package validators

import (
	"reflect"
	"testing"

	"github.com/lccezinha/twodo/internal/twodo"
)

func TestTodoValidator(t *testing.T) {
	t.Run("When todo title is blank, it returns errs", func(t *testing.T) {
		validator := TodoValidator{}
		todo := twodo.Todo{Title: ""}
		expectedErrs := []twodo.ValidationError{
			twodo.ValidationError{
				Field:   "Title",
				Message: "Can not be blank",
				Type:    "Required",
			},
		}

		errs := validator.Validate(todo)

		if !reflect.DeepEqual(expectedErrs, errs) {
			t.Errorf("Expected %v - Received: %v", expectedErrs, errs)
		}
	})
}
