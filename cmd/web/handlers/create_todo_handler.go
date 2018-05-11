package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lccezinha/twodo/cmd/web/presenters"
	"github.com/lccezinha/twodo/internal/twodo"
)

type todoParams struct {
	Description string
}

// CreateTodoHandler struct represents create handler
type CreateTodoHandler struct {
	UseCase          twodo.CreateUseCase
	PresenterFactory presenters.PresenterFactory
}

func (c *CreateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	presenter := c.PresenterFactory.Create(w)
	var params todoParams
	json.NewDecoder(r.Body).Decode(&params)
	c.UseCase.Run(params.Description, presenter)
}

// NewCreateTodoHandler works as a factory method
func NewCreateTodoHandler(u twodo.CreateUseCase, pf presenters.PresenterFactory) *CreateTodoHandler {
	return &CreateTodoHandler{UseCase: u, PresenterFactory: pf}
}
