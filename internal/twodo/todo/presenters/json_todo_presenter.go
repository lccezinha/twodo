package presenters

import (
	"net/http"

	"github.com/lccezinha/twodo/internal/twodo"
)

type JSONTodoPresenter struct {
	ResponseWriter http.ResponseWriter
}

func (tp *JSONTodoPresenter) Present(status int, todo twodo.Todo) {

}
