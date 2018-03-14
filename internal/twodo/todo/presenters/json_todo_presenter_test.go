package presenters

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/lccezinha/twodo/internal/twodo"
)

func TestPresentTodo(t *testing.T) {
	w := httptest.NewRecorder()
	presenter := JSONTodoPresenter{w}
	todo := twodo.Todo{
		ID:          1,
		Title:       "Title",
		Description: "Description",
		Done:        false,
	}
	expectedStatus := http.StatusOK

	presenter.Present(http.StatusOK, todo)

	expectedBody := []byte(
		fmt.Sprintf(
			`{"id":%d,"title":"%s", "description":"%s", "done":"%v"}`, todo.ID, todo.Title, todo.Description, todo.Done,
		),
	)

	response := w.Result()
	body, _ := ioutil.ReadAll(response.Body)

	if !reflect.DeepEqual(body, expectedBody) {
		t.Errorf("Expected: %s. Actual: %s", expectedBody, body)
	}

	if w.Result().StatusCode != expectedStatus {
		t.Errorf("Expected: %d. Actual: %d", expectedStatus, w.Result().StatusCode)
	}

	expectedContentType := "application/json"
	contentType := w.Result().Header.Get("Content-Type")

	if contentType != expectedContentType {
		t.Errorf("Expected: %s. Actual: %s", expectedContentType, contentType)
	}
}
