package todo

import (
	"github.com/lccezinha/twodo/internal/twodo"
	"github.com/lccezinha/twodo/internal/twodo/todo/validators"
)

// CreateService define a service to create a new Todo
type CreateService struct {
	repository twodo.Repository
	validator  twodo.Validator
}

// Run method will execute the action of create a new Todo
func (c *CreateService) Run(title string, description string) []twodo.ValidationError {
	todo := twodo.Todo{
		Title:       title,
		Description: description,
		// CreatedAt:   time.Now(),
		Done: false,
	}

	errs := c.validator.Validate(todo)
	if len(errs) > 0 {
		return errs
	}

	return nil
}

// NewCreateService works as a factory method
func NewCreateService(r twodo.Repository) *CreateService {
	return &CreateService{
		repository: r,
		validator:  validators.TodoValidator{},
	}
}
