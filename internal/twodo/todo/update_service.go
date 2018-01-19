package todo

import "github.com/lccezinha/twodo/internal/twodo"

// UpdateService define a service to list all todos
type UpdateService struct {
	repository twodo.Repository
}

// Run method will execute the action of update a Todo
func (u *UpdateService) Run(id int) error {
	err := u.repository.Update(id)

	if err != nil {
		return err
	}

	return nil
}

// NewUpdateService works as a factory method
func NewUpdateService(r twodo.Repository) *UpdateService {
	return &UpdateService{r}
}
