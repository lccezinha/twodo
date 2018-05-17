package handlers

import (
	"context"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/lccezinha/twodo/internal/test/fakes"
)

func TestMarkAsUndoneHandler(t *testing.T) {
	repository := fakes.NewFakeRepository()
	usecase := fakes.NewFakeMarkAsUndoneUseCase(repository)
	presenterFactory := fakes.NewFakePresenterFactory()
	handler := MarkAsUndoneHandler{
		UseCase:          usecase,
		PresenterFactory: presenterFactory,
	}

	response := httptest.NewRecorder()

	request := httptest.NewRequest("PUT", "http://localhost:8888/api/v1/todos/1/mark-as-undone", nil)
	params := httprouter.Params{httprouter.Param{Key: "id", Value: strconv.Itoa(1)}}
	request = request.WithContext(context.WithValue(request.Context(), httprouter.ParamsKey, params))

	handler.ServeHTTP(response, request)

	if presenterFactory.ResponseWriter != response {
		t.Errorf("Expected: %v. Actual: %v", response, presenterFactory.ResponseWriter)
	}
}
