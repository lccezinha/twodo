package usecases

import "github.com/lccezinha/twodo/internal/twodo"

type MarkAsUndoneUseCase struct {
	repository twodo.Repository
}

func (m *MarkAsUndoneUseCase) Run(id int, presenter twodo.Presenter) error {
	todo, err := m.repository.Update(id, false)

	if err != nil {
		panic(err)
	}

	presenter.PresentUpdated(todo)

	return nil
}

func NewMarkAsUndoneUseCase(r twodo.Repository) *MarkAsUndoneUseCase {
	return &MarkAsUndoneUseCase{repository: r}
}
