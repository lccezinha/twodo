package handlers

import (
	"net/http"

	"github.com/lccezinha/twodo/cmd/web/presenters"
	"github.com/lccezinha/twodo/internal/twodo"
)

type ListAllTodosHandler struct {
	UseCase          twodo.ListAllUseCase
	PresenterFactory presenters.PresenterFactory
}

func (l *ListAllTodosHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	presenter := l.PresenterFactory.Create(w)
	l.UseCase.Run(presenter)
}

func NewListAllTodosHandler(u twodo.ListAllUseCase, pf presenters.PresenterFactory) *ListAllTodosHandler {
	return &ListAllTodosHandler{
		UseCase:          u,
		PresenterFactory: pf,
	}
}
