package usecases

import (
	"github.com/lccezinha/twodo/internal/twodo"
	"github.com/lccezinha/twodo/internal/twodo/todo/validators"
)

// CreateTodoUseCase define a service to create a new Todo
type CreateTodoUseCase struct {
	repository twodo.Repository
	validator  twodo.Validator
}

// Run method will execute the action of create a new Todor
func (c *CreateTodoUseCase) Run(description string, presenter twodo.Presenter) error {
	todo := twodo.Todo{
		Description: description,
		// CreatedAt:   time.Now(),
		Done: false,
	}

	errs := c.validator.Validate(todo)
	if len(errs) > 0 {
		presenter.PresentErrors(errs)
		return nil
	}

	todo, err := c.repository.Save(todo)
	if err != nil {
		return err
	}
	presenter.PresentCreatedTodo(todo)

	return nil
}

// NewCreateTodoUseCase works as a factory method
func NewCreateTodoUseCase(r twodo.Repository) *CreateTodoUseCase {
	return &CreateTodoUseCase{
		repository: r,
		validator:  validators.TodoValidator{},
	}
}
