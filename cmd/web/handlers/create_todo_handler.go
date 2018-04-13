package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lccezinha/twodo/internal/twodo"
)

type todoParams struct {
	Description string
	Title       string
}

// CreateTodoHandler struct represents create handler
type CreateTodoHandler struct {
	UseCase   twodo.CreateUseCase
	Presenter twodo.Presenter
}

func (c *CreateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var params todoParams
	json.NewDecoder(r.Body).Decode(&params)

	err := c.UseCase.Run(params.Title, params.Description, c.Presenter)
	if err != nil {
		// c.Presenter.PresentErrors(errs)
	}
}

// NewCreateTodoHandler works as a factory method
func NewCreateTodoHandler(u twodo.CreateUseCase, p twodo.Presenter) *CreateTodoHandler {
	return &CreateTodoHandler{UseCase: u, Presenter: p}
}
