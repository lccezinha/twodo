package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/lccezinha/twodo/internal/twodo"
)

// JSONTodoPresenter will handle a Todo as JSON
type JSONTodoPresenter struct {
	ResponseWriter http.ResponseWriter
}

// Present will present a Todo as a JSON
func (p *JSONTodoPresenter) Present(status int, todo twodo.Todo) {
	todoJSON, _ := json.Marshal(todo)
	p.ResponseWriter.Header().Set("Content-type", "application/json")
	p.ResponseWriter.WriteHeader(status)
	p.ResponseWriter.Write([]byte(todoJSON))
}

// PresentErrors will present all errors for a Todo as a JSON
func (p *JSONTodoPresenter) PresentErrors(errs []twodo.ValidationError) {
	errsJSON, _ := json.Marshal(showErrs{Errors: errs})
	p.ResponseWriter.Header().Set("Content-type", "application/json")
	p.ResponseWriter.WriteHeader(http.StatusBadRequest)
	p.ResponseWriter.Write([]byte(errsJSON))
}

type showErrs struct {
	Errors []twodo.ValidationError `json:"errors"`
}
