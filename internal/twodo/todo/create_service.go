package todo

import (
	"errors"
	"time"

	"github.com/lccezinha/twodo/internal/twodo"
)

var errCannotBeBlank = errors.New("can not be blank")

// CreateService define a service to create a new Todo
type CreateService struct {
	repository twodo.Repository
}

// Run method will execute the action of create a new Todo
func (c *CreateService) Run(title string, description string) error {
	if title == "" {
		return errCannotBeBlank
	}

	todo := twodo.Todo{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		Done:        false,
	}

	if _, err := c.repository.Save(todo); err != nil {
		return err
	}

	return nil
}

// NewCreateService works as a factory method
func NewCreateService(r twodo.Repository) twodo.Creator {
	return &CreateService{r}
}
