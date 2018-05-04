package usecases

import "github.com/lccezinha/twodo/internal/twodo"

// ListAllUseCase define a service to list all todos
type ListAllUseCase struct {
	repository twodo.Repository
}

func (l *ListAllUseCase) Run(presenter twodo.Presenter) ([]twodo.Todo, error) {
	todos, err := l.repository.ListAll()

	if err != nil {
		panic(err.Error())
	}

	presenter.PresentListTodos(todos)

	return nil, nil
}

func NewListAllUseCase(r twodo.Repository) *ListAllUseCase {
	return &ListAllUseCase{repository: r}
}
