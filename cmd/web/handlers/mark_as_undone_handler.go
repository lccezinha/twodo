package handlers

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/lccezinha/twodo/cmd/web/presenters"
	"github.com/lccezinha/twodo/internal/twodo"
)

type MarkAsUndoneHandler struct {
	UseCase          twodo.MarkAsDoneUseCase
	PresenterFactory presenters.PresenterFactory
}

func (m *MarkAsUndoneHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	presenter := m.PresenterFactory.Create(w)
	params := httprouter.ParamsFromContext(r.Context())
	id, _ := strconv.Atoi(params.ByName("id"))
	m.UseCase.Run(id, presenter)
}

func NewMarkAsUndoneHandler(u twodo.MarkAsDoneUseCase, pf presenters.PresenterFactory) *MarkAsUndoneHandler {
	return &MarkAsUndoneHandler{
		UseCase:          u,
		PresenterFactory: pf,
	}
}
