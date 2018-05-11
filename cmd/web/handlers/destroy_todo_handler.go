package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/lccezinha/twodo/cmd/web/presenters"
	"github.com/lccezinha/twodo/internal/twodo"
)

// DestroyTodoHandler struct represents create handler
type DestroyTodoHandler struct {
	UseCase          twodo.DestroyUseCase
	PresenterFactory presenters.PresenterFactory
}

func (d *DestroyTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	presenter := d.PresenterFactory.Create(w)
	path := strings.Split(r.URL.Path, "/")
	idString := path[len(path)-1]
	id, _ := strconv.Atoi(idString)

	d.UseCase.Run(id, presenter)
}

// NewDestroyTodoHandler works as a factory method
func NewDestroyTodoHandler(u twodo.DestroyUseCase, pf presenters.PresenterFactory) *DestroyTodoHandler {
	return &DestroyTodoHandler{
		UseCase:          u,
		PresenterFactory: pf,
	}
}
