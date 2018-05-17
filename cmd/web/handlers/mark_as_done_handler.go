package handlers

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/lccezinha/twodo/cmd/web/presenters"
	"github.com/lccezinha/twodo/internal/twodo"
)

type MarkAsDoneHandler struct {
	UseCase          twodo.MarkAsDoneUseCase
	PresenterFactory presenters.PresenterFactory
}

func (m *MarkAsDoneHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	presenter := m.PresenterFactory.Create(w)
	params := httprouter.ParamsFromContext(r.Context())
	id, _ := strconv.Atoi(params.ByName("id"))
	m.UseCase.Run(id, presenter)
}

func NewMarkAsDoneHandler(u twodo.MarkAsDoneUseCase, pf presenters.PresenterFactory) *MarkAsDoneHandler {
	return &MarkAsDoneHandler{
		UseCase:          u,
		PresenterFactory: pf,
	}
}
