package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/lccezinha/twodo/internal/test/fakes"
)

func TestDestroyTodoHandler(t *testing.T) {
	t.Run("Given a valid id remove it", func(t *testing.T) {
		repository := fakes.NewFakeRepository()
		usecase := fakes.NewDestroyUseCase(repository)
		presenterFactory := fakes.NewFakePresenterFactory()
		handler := DestroyTodoHandler{
			UseCase:          usecase,
			PresenterFactory: presenterFactory,
		}

		response := httptest.NewRecorder()
		request := httptest.NewRequest("DELETE", "http://localhost:8000/api/v1/todos/1", nil)

		handler.ServeHTTP(response, request)

		if !presenterFactory.CreateCalled {
			t.Error("PresentFactory not called")
		}

		if presenterFactory.ResponseWriter != response {
			t.Errorf("Expected: %v. Actual: %v", response, presenterFactory.ResponseWriter)
		}
	})
}
