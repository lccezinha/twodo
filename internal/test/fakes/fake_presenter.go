package fakes

import (
	"net/http"

	"github.com/lccezinha/twodo/internal/twodo"
)

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

func NewFakePresenter() *FakePresenter {
	return &FakePresenter{}
}

type FakePresenterFactory struct {
	CreateCalled   bool
	ResponseWriter http.ResponseWriter
	Presenter      *FakePresenter
}

func (pf *FakePresenterFactory) Create(w http.ResponseWriter) twodo.Presenter {
	pf.CreateCalled = true
	pf.ResponseWriter = w
	return pf.Presenter
}

func NewFakePresenterFactory() *FakePresenterFactory {
	return &FakePresenterFactory{}
}
