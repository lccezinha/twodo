package application

import (
	"github.com/lccezinha/twodo/cmd/env"
	"github.com/lccezinha/twodo/cmd/web/handlers"
	"github.com/lccezinha/twodo/cmd/web/presenters"
	"github.com/lccezinha/twodo/internal/twodo/todo/usecases"
)

// Handlers is a simple struct that hold all app handlers
type Handlers struct {
	CreateTodo   *handlers.CreateTodoHandler
	ListAllTodos *handlers.ListAllTodosHandler
	DestroyTodo  *handlers.DestroyTodoHandler
	MarkAsDone   *handlers.MarkAsDoneHandler
	MarkAsUndone *handlers.MarkAsUndoneHandler
}

// InitializeHandlers will return all handler initialized
func InitializeHandlers() *Handlers {
	app := env.Init()
	createUseCase := usecases.NewCreateTodoUseCase(app.Repository)
	listAllUseCase := usecases.NewListAllUseCase(app.Repository)
	destroyUseCase := usecases.NewDestroyUseCase(app.Repository)
	markAsDoneUseCase := usecases.NewMarkAsDoneUseCase(app.Repository)
	markAsUndoneUseCase := usecases.NewMarkAsUndoneUseCase(app.Repository)
	presenterFactory := presenters.NewPresenterFactory()

	return &Handlers{
		CreateTodo:   handlers.NewCreateTodoHandler(createUseCase, presenterFactory),
		ListAllTodos: handlers.NewListAllTodosHandler(listAllUseCase, presenterFactory),
		DestroyTodo:  handlers.NewDestroyTodoHandler(destroyUseCase, presenterFactory),
		MarkAsDone:   handlers.NewMarkAsDoneHandler(markAsDoneUseCase, presenterFactory),
		MarkAsUndone: handlers.NewMarkAsUndoneHandler(markAsUndoneUseCase, presenterFactory),
	}
}
