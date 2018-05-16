package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/lccezinha/twodo/internal/test/fakes"
)

func TestMarkAsDoneHandler(t *testing.T) {
	repository := fakes.NewFakeRepository()
	usecase := fakes.NewFakeMarkAsDoneUseCase(repository)
	presenterFactory := fakes.NewFakePresenterFactory()
	handler := MarkAsDoneHandler{
		UseCase:          usecase,
		PresenterFactory: presenterFactory,
	}

	response := httptest.NewRecorder()
	request := httptest.NewRequest("PUT", "http://localhost:8888/api/v1/todos/1", nil)

	handler.ServeHTTP(response, request)

	if presenterFactory.ResponseWriter != response {
		t.Errorf("Expected: %v. Actual: %v", response, presenterFactory.ResponseWriter)
	}
}
