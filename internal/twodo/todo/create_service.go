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
	if title == "" || description == "" {
		return errCannotBeBlank
	}

	t := &twodo.Todo{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}

	if err := c.repository.Save(t); err != nil {
		return err
	}

	return nil
}

// NewCreateService works as a factory method
func NewCreateService(r twodo.Repository) *CreateService {
	return &CreateService{r}
}
