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
	p.ResponseWriter.Write([]byte(todoJSON))
	p.ResponseWriter.WriteHeader(status)
}

// PresentErrors will present all errors for a Todo as a JSON
func (p *JSONTodoPresenter) PresentErrors(errs []twodo.ValidationError) {
	// true
}
