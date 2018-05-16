package usecases

import "github.com/lccezinha/twodo/internal/twodo"

type MarkAsDoneUseCase struct {
	repository twodo.Repository
}

func (m *MarkAsDoneUseCase) Run(id int, presenter twodo.Presenter) error {
	todo, err := m.repository.Update(id, true)

	if err != nil {
		panic(err)
	}

	presenter.PresentUpdated(todo)

	return nil
}

func NewMarkAsDoneUseCase(r twodo.Repository) *MarkAsDoneUseCase {
	return &MarkAsDoneUseCase{repository: r}
}
