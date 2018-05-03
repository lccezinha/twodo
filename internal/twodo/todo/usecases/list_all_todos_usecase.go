package usecases

import "github.com/lccezinha/twodo/internal/twodo"

type ListAllUseCase struct {
	repository twodo.Repository
}

func (l *ListAllUseCase) Run(presenter twodo.Presenter) ([]twodo.Todo, error) {
	return nil, nil
}

func NewListAllUseCase(r twodo.Repository) *ListAllUseCase {
	return &ListAllUseCase{repository: r}
}
