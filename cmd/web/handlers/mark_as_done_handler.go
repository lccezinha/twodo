package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/lccezinha/twodo/cmd/web/presenters"
	"github.com/lccezinha/twodo/internal/twodo"
)

type MarkAsDoneHandler struct {
	UseCase          twodo.MarkAsDoneUseCase
	PresenterFactory presenters.PresenterFactory
}

func (m *MarkAsDoneHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	presenter := m.PresenterFactory.Create(w)
	path := strings.Split(r.URL.Path, "/")
	idString := path[len(path)-1]
	id, _ := strconv.Atoi(idString)
	m.UseCase.Run(id, presenter)
}

func NewMarkAsDoneHandler(u twodo.MarkAsDoneUseCase, pf presenters.PresenterFactory) *MarkAsDoneHandler {
	return &MarkAsDoneHandler{
		UseCase:          u,
		PresenterFactory: pf,
	}
}
