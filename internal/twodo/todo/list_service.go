package todo

import "github.com/lccezinha/twodo/internal/twodo"

// ListService define a service to list all todos
type ListService struct {
	repository twodo.Repository
}

// Run method will execute the action of list all Todos
func (l *ListService) Run() ([]*twodo.Todo, error) {
	todos, err := l.repository.ListAll()

	return todos, err
}

// NewListService works as a factory method
func NewListService(r twodo.Repository) *ListService {
	return &ListService{r}
}
