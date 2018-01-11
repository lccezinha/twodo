package todo

import (
	"github.com/lccezinha/twodo/internal/twodo"
)

// DestroyService define a service to destroy a Todo
type DestroyService struct {
	repository twodo.Repository
}

// Run method will execute the action of destroy Todo
func (d *DestroyService) Run(id int) error {
	result := d.repository.Destroy(id)

	return result
}

// NewDestroyService works as a factory method
func NewDestroyService(r twodo.Repository) *DestroyService {
	return &DestroyService{r}
}
