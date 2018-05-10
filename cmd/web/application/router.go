package application

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// InitializeRouter will return all handler
func InitializeRouter() http.Handler {
	handlers := InitializeHandlers()

	router := httprouter.New()
	router.Handler(http.MethodGet, "/api/v1/todos", handlers.ListAllTodos)
	router.Handler(http.MethodPost, "/api/v1/todos", handlers.CreateTodo)

	return router
}
