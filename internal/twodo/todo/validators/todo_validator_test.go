package validators

import (
	"reflect"
	"testing"

	"github.com/lccezinha/twodo/internal/twodo"
)

func TestTodoValidator(t *testing.T) {
	t.Run("When todo description is blank, it returns errs", func(t *testing.T) {
		validator := TodoValidator{}
		todo := twodo.Todo{Description: ""}
		expectedErrs := []twodo.ValidationError{
			twodo.ValidationError{
				Field:   "Description",
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
