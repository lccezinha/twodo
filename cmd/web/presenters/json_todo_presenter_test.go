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

func TestPresentCreatedTodo(t *testing.T) {
	w := httptest.NewRecorder()
	presenter := JSONTodoPresenter{w}
	todo := twodo.Todo{
		ID:          1,
		Description: "Description",
		// CreatedAt:   time.Now(),
		Done: false,
	}
	expectedBody := []byte(
		fmt.Sprintf(
			`{"id":%d,"description":"%s","done":%t}`,
			todo.ID, todo.Description, todo.Done,
		),
	)

	presenter.PresentCreatedTodo(todo)

	response := w.Result()
	contentType := response.Header.Get("Content-Type")
	body, _ := ioutil.ReadAll(response.Body)

	expectedStatus := http.StatusCreated
	expectedContentType := "application/json"

	if !reflect.DeepEqual(body, expectedBody) {
		t.Errorf("Expected: %s. Actual: %s", expectedBody, body)
	}

	if response.StatusCode != expectedStatus {
		t.Errorf("Expected: %d. Actual: %d", expectedStatus, response.StatusCode)
	}

	if contentType != expectedContentType {
		t.Errorf("Expected: %s. Actual: %s", expectedContentType, contentType)
	}
}

func TestPresentTodoErrs(t *testing.T) {
	w := httptest.NewRecorder()
	presenter := JSONTodoPresenter{w}
	errs := []twodo.ValidationError{
		twodo.ValidationError{
			Field:   "Title",
			Message: "Can not be blank",
			Type:    "Required",
		},
	}

	presenter.PresentErrors(errs)

	expectedError := fmt.Sprintf(`{"field":"%s","message":"%s","type":"%s"}`, errs[0].Field, errs[0].Message, errs[0].Type)
	expectedBody := []byte(`{"errors":[` + expectedError + `]}`)
	expectedContentType := "application/json"
	expectedStatus := http.StatusBadRequest

	response := w.Result()
	body, _ := ioutil.ReadAll(response.Body)
	contentType := w.Result().Header.Get("Content-Type")

	if !reflect.DeepEqual(body, expectedBody) {
		t.Errorf("Expected: %s. Actual: %s", expectedBody, body)
	}

	if w.Result().StatusCode != expectedStatus {
		t.Errorf("Expected: %d. Actual: %d", expectedStatus, w.Result().StatusCode)
	}

	if contentType != expectedContentType {
		t.Errorf("Expected: %s. Actual: %s", expectedContentType, contentType)
	}
}

func TestPresentInvalidHTTPMethodError(t *testing.T) {
	w := httptest.NewRecorder()
	presenter := JSONTodoPresenter{w}
	presenter.PresentInvalidHTTPMethodError("POST")
	response := w.Result()
	expectedErrorMessage := "Wrong HTTP method"
	expectedHTTPMethod := "POST"

	if response.Header.Get("allowed") != "POST" {
		t.Errorf("Expected: %s. Actual: %s", expectedHTTPMethod, response.Header.Get("allowed"))
	}

	if response.Header.Get("error_message") != expectedErrorMessage {
		t.Errorf("Expected: %s. Actual: %s", expectedErrorMessage, response.Header.Get("error_message"))
	}

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected: %d. Actual: %d", http.StatusMethodNotAllowed, response.StatusCode)
	}
}

func TestPresentListTodos(t *testing.T) {
	t.Run("Return presenter with empty list", func(t *testing.T) {
		w := httptest.NewRecorder()
		presenter := JSONTodoPresenter{w}
		todos := []twodo.Todo{}
		presenter.PresentListTodos(todos)
		expectedStatus := http.StatusOK
		response := w.Result()
		body, _ := ioutil.ReadAll(response.Body)
		expectedBody := []byte(fmt.Sprint(`[]`))

		if response.StatusCode != expectedStatus {
			t.Errorf("Expected: %d. Actual: %d", expectedStatus, response.StatusCode)
		}

		if !reflect.DeepEqual(body, expectedBody) {
			t.Errorf("Expected: %s. Actual: %s", expectedBody, body)
		}
	})

	t.Run("Return presenter with filled list", func(t *testing.T) {
		w := httptest.NewRecorder()
		presenter := JSONTodoPresenter{w}
		todoOne := twodo.Todo{ID: 1, Description: "Some Todo #1", Done: true}
		todoTwo := twodo.Todo{ID: 2, Description: "Some Todo #2", Done: false}
		todos := []twodo.Todo{todoOne, todoTwo}

		presenter.PresentListTodos(todos)

		expectedStatus := http.StatusOK
		response := w.Result()
		body, _ := ioutil.ReadAll(response.Body)
		expectedBody := []byte(
			fmt.Sprintf(
				`[{"id":%d,"description":"%s","done":%t},{"id":%d,"description":"%s","done":%t}]`,
				todoOne.ID, todoOne.Description, todoOne.Done,
				todoTwo.ID, todoTwo.Description, todoTwo.Done,
			),
		)

		if response.StatusCode != expectedStatus {
			t.Errorf("Expected: %d. Actual: %d", expectedStatus, response.StatusCode)
		}

		if !reflect.DeepEqual(body, expectedBody) {
			t.Errorf("Expected: %s. Actual: %s", expectedBody, body)
		}
	})
}
