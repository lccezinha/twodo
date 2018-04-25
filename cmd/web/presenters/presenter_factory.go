package presenters

import (
	"net/http"

	"github.com/lccezinha/twodo/internal/twodo"
)

// PresenterFactory represents a factory
type PresenterFactory interface {
	Create(http.ResponseWriter) twodo.Presenter
}
