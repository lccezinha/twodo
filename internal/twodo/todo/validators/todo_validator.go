package validators

import "github.com/lccezinha/twodo/internal/twodo"

// TodoValidator is a simple struct to hold the Validate method
type TodoValidator struct{}

// Validate will handle all validations for a Todo
func (v TodoValidator) Validate(todo twodo.Todo) []twodo.ValidationError {
	errs := []twodo.ValidationError{}

	if todo.Title == "" {
		err := twodo.ValidationError{
			Field:   "Title",
			Message: "Can not be blank",
			Type:    "Required",
		}

		errs = append(errs, err)
	}

	return errs
}
