package usecases

import "github.com/lccezinha/twodo/internal/twodo"

type DestroyUseCase struct {
	repository twodo.Repository
}

func (d *DestroyUseCase) Run(id int, presenter twodo.Presenter) error {
	err := d.repository.Destroy(id)

	if err != nil {
		panic(err)
	}

	presenter.PresentDestroyed()

	return nil
}

func NewDestroyUseCase(r twodo.Repository) *DestroyUseCase {
	return &DestroyUseCase{repository: r}
}
