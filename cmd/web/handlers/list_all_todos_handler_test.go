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

	response := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://localhost:8888/api/v1/todos", nil)

	handler.ServeHTTP(response, request)

	if presenterFactory.ResponseWriter != response {
		t.Errorf("Expected: %v. Actual: %v", response, presenterFactory.ResponseWriter)
	}
}
