package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/lccezinha/twodo/internal/test/fakes"
)

func TestListAllTodos(t *testing.T) {
	repository := fakes.NewFakeRepository()
	usecase := fakes.NewFakeListUseCase(repository)
	presenterFactory := fakes.NewFakePresenterFactory()
	handler := ListAllTodosHandler{
		UseCase:          usecase,
		PresenterFactory: presenterFactory,
	}

	// todos := []twodo.Todo{
	// 	twodo.Todo{ID: 1, Description: "Todo #1"},
	// 	twodo.Todo{ID: 2, Description: "Todo #2"},
	// }

	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://localhost:8888/api/v1/todos", nil)

	handler.ServeHTTP(response, request)

	if presenterFactory.ResponseWriter != response {
		t.Errorf("Expected: %v. Actual: %v", response, presenterFactory.ResponseWriter)
	}
}
