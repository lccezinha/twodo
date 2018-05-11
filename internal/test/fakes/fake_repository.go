package fakes

import "github.com/lccezinha/twodo/internal/twodo"

type FakeRepository struct {
	Todo          twodo.Todo
	ListAllResult []twodo.Todo
}

func (fr *FakeRepository) Save(t twodo.Todo) (twodo.Todo, error) {
	fr.Todo = t

	return fr.Todo, nil
}

func (fr *FakeRepository) ListAll() ([]twodo.Todo, error) {
	return fr.ListAllResult, nil
}

func (fr *FakeRepository) Destroy(id int) error {
	return nil
}

func NewFakeRepository() *FakeRepository {
	return &FakeRepository{}
}
