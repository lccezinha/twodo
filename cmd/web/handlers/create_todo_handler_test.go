package handlers

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lccezinha/twodo/internal/test/fakes"
)

func TestCreateTodoHandler(t *testing.T) {
	t.Run("Given valid args, call use case with args", func(t *testing.T) {
		repository := fakes.NewFakeRepository()
		usecase := fakes.NewFakeUseCase(repository)
		presenterFactory := fakes.NewFakePresenterFactory()
		handler := CreateTodoHandler{
			UseCase:          usecase,
			PresenterFactory: presenterFactory,
		}

		description := "Description"
		params := `{"description": "` + description + `"}`

		response := httptest.NewRecorder()
		request := httptest.NewRequest("POST", "http://localhost:8080/api/v1/todos", strings.NewReader(params))

		handler.ServeHTTP(response, request)

		if presenterFactory.ResponseWriter != response {
			t.Errorf("Expected: %v. Actual: %v", response, presenterFactory.ResponseWriter)
		}

		if usecase.Description != description {
			t.Errorf("Expected: %v. Actual: %v", description, usecase.Description)
		}
	})
}
