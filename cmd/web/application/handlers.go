package application

import (
	"github.com/lccezinha/twodo/cmd/repository/todo"
	"github.com/lccezinha/twodo/cmd/web/handlers"
	"github.com/lccezinha/twodo/cmd/web/presenters"
	"github.com/lccezinha/twodo/internal/twodo/todo/usecases"
)

// Handlers is a simple struct that hold all app handlers
type Handlers struct {
	CreateTodo   *handlers.CreateTodoHandler
	ListAllTodos *handlers.ListAllTodosHandler
}

// InitializeHandlers will return all handler initialized
func InitializeHandlers() *Handlers {
	repository := todo.NewMemoryRepository()
	createUseCase := usecases.NewCreateTodoUseCase(repository)
	listAllUseCase := usecases.NewListAllUseCase(repository)
	presenterFactory := presenters.NewPresenterFactory()

	return &Handlers{
		CreateTodo:   handlers.NewCreateTodoHandler(createUseCase, presenterFactory),
		ListAllTodos: handlers.NewListAllTodosHandler(listAllUseCase, presenterFactory),
	}
}
