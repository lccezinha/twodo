package stubs

import "github.com/lccezinha/twodo/internal/twodo"

type FakePresenter struct {
	Todo twodo.Todo
	Errs []twodo.ValidationError
}

func (fp *FakePresenter) PresentCreatedTodo(todo twodo.Todo) {
	fp.Todo = todo
}

func (fp *FakePresenter) PresentErrors(errs []twodo.ValidationError) {
	fp.Errs = errs
}

type FakeRepository struct {
	Todo twodo.Todo
}

func NewFakePresenter() *FakePresenter {
	return new(FakePresenter)
}

func (fr *FakeRepository) Save(t twodo.Todo) (twodo.Todo, error) {
	fr.Todo = t

	return fr.Todo, nil
}

func NewFakeRepository() *FakeRepository {
	return new(FakeRepository)
}
