package fakes

import "github.com/lccezinha/twodo/internal/twodo"

type FakeUseCase struct {
	Repository  twodo.Repository
	Title       string
	Description string
	Presenter   twodo.Presenter
}

func (fu *FakeUseCase) Run(title string, description string, presenter twodo.Presenter) error {
	fu.Title = title
	fu.Description = description
	fu.Presenter = presenter
	return nil
}

func NewFakeUseCase(r twodo.Repository) *FakeUseCase {
	return &FakeUseCase{Repository: r}
}
