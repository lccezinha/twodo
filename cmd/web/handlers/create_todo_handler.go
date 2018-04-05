package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lccezinha/twodo/internal/twodo"
)

type todoParams struct {
	Todo struct {
		Title       string
		Description string
	}
}

type CreateTodoHandler struct {
	UseCase   twodo.CreateUseCase
	Presenter twodo.Presenter
}

func (c *CreateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := c.extractParams(r)

	err := c.UseCase.Run(params.Todo.Title, params.Todo.Description, c.Presenter)
	if err != nil {
		log.Print(err)
		// c.Presenter.PresentErrors(errs)
	}
}

// NewCreateTodoHandler works as a factory method
func NewCreateTodoHandler(u twodo.CreateUseCase, p twodo.Presenter) *CreateTodoHandler {
	return &CreateTodoHandler{UseCase: u, Presenter: p}
}

func (c *CreateTodoHandler) extractParams(r *http.Request) todoParams {
	var params todoParams
	json.NewDecoder(r.Body).Decode(&params)
	return params
}
