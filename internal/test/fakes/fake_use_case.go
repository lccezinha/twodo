package fakes

import "github.com/lccezinha/twodo/internal/twodo"

type FakeCreateUseCase struct {
	Repository  twodo.Repository
	Description string
	Presenter   twodo.Presenter
}

func (fu *FakeCreateUseCase) Run(description string, presenter twodo.Presenter) error {
	fu.Description = description
	fu.Presenter = presenter
	return nil
}

func NewFakeCreateUseCase(r twodo.Repository) *FakeCreateUseCase {
	return &FakeCreateUseCase{Repository: r}
}

type FakeListUseCase struct {
	Repository twodo.Repository
	List       []twodo.Todo
	Presenter  twodo.Presenter
}

func (fu *FakeListUseCase) Run(presenter twodo.Presenter) ([]twodo.Todo, error) {
	return []twodo.Todo{}, nil
}

func NewFakeListUseCase(r twodo.Repository) *FakeListUseCase {
	return &FakeListUseCase{Repository: r}
}
