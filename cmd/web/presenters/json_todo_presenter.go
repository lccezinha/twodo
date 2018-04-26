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

// PresentCreatedTodo will present a Todo as a JSON
func (p *JSONTodoPresenter) PresentCreatedTodo(todo twodo.Todo) {
	todoJSON, _ := json.Marshal(todo)
	p.renderJSONResponse(http.StatusCreated, todoJSON)
}

// PresentErrors will present all errors for a Todo as a JSON
func (p *JSONTodoPresenter) PresentErrors(errs []twodo.ValidationError) {
	errsJSON, _ := json.Marshal(showErrs{Errors: errs})
	p.renderJSONResponse(http.StatusBadRequest, errsJSON)
}

// PresentInvalidHTTPMethodError response when http is called using wrong method
func (p *JSONTodoPresenter) PresentInvalidHTTPMethodError(am string) {
	p.ResponseWriter.Header().Set("allowed", am)
	p.ResponseWriter.Header().Set("error_message", "Wrong HTTP method")
	p.ResponseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *JSONTodoPresenter) renderJSONResponse(status int, data []byte) {
	p.ResponseWriter.Header().Set("Content-type", "application/json")
	p.ResponseWriter.WriteHeader(status)
	p.ResponseWriter.Write([]byte(data))
}

type showErrs struct {
	Errors []twodo.ValidationError `json:"errors"`
}

// Create implement PresenterFactory
func (p JSONTodoPresenter) Create(w http.ResponseWriter) twodo.Presenter {
	return &JSONTodoPresenter{w}
}

func NewPresenterFactory() PresenterFactory {
	return JSONTodoPresenter{}
}
