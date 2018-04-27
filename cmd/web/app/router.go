package app

import "net/http"

// InitializeRouter will return all handler
func InitializeRouter() http.Handler {
	handlers := InitializeHandlers()

	router := http.NewServeMux()
	router.Handle("/api/v1/todos", handlers.CreateTodo)

	return router
}
