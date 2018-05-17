package handlers

import (
	"context"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/julienschmidt/httprouter"
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

	request := httptest.NewRequest("PUT", "http://localhost:8888/api/v1/todos/1/markasdone", nil)
	params := httprouter.Params{httprouter.Param{Key: "id", Value: strconv.Itoa(1)}}
	request = request.WithContext(context.WithValue(request.Context(), httprouter.ParamsKey, params))

	handler.ServeHTTP(response, request)

	if presenterFactory.ResponseWriter != response {
		t.Errorf("Expected: %v. Actual: %v", response, presenterFactory.ResponseWriter)
	}
}
