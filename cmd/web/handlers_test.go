package web

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	responseRecorder := httptest.NewRecorder()
	IndexHandler(responseRecorder, req)

	response := responseRecorder.Result()

	if status := response.StatusCode; status != http.StatusOK {
		t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusOK, response.StatusCode)
	}

	body, _ := ioutil.ReadAll(response.Body)

	if !strings.Contains(string(body), "Twodo APP") {
		t.Error("Body does not contain correct app title")
	}
}

func TestCreateHandler(t *testing.T) {
	t.Run("Success create a new todo", func(t *testing.T) {
		params := url.Values{}
		params.Add("title", "Some title")
		params.Add("description", "Some Description")
		req := httptest.NewRequest("POST", "/todos", strings.NewReader(params.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateHandler)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusFound {
			t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusFound, w.Code)
		}
	})

	t.Run("Fail - use wrong http method to send request", func(t *testing.T) {
		params := url.Values{}
		params.Add("title", "")
		params.Add("description", "")
		req := httptest.NewRequest("GET", "/todos", strings.NewReader(params.Encode()))
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateHandler)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("Fail because todo title is not set", func(t *testing.T) {
		params := url.Values{}
		params.Add("title", "")
		params.Add("description", "Some Description")
		req := httptest.NewRequest("POST", "/todos", strings.NewReader(params.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateHandler)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusBadRequest, w.Code)
		}
	})
}

func TestDestroyHandler(t *testing.T) {
	t.Run("Success when resource is destroyed", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/destroy?todoID=1", nil)
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(DestroyHandler)
		handler.ServeHTTP(response, req)

		if status := response.Code; status != http.StatusNoContent {
			t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusNoContent, response.Code)
		}
	})

	t.Run("Fail - use wrong http method to send request", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/destroy?todoID=1", nil)
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(DestroyHandler)
		handler.ServeHTTP(response, req)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusBadRequest, response.Code)
		}
	})
}

func TestDoneHandler(t *testing.T) {
	t.Run("Success when update a resource to true", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/done?todoID=1", nil)
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(DoneHandler)
		handler.ServeHTTP(response, req)

		if status := response.Code; status != http.StatusOK {
			t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusOK, response.Code)
		}
	})

	t.Run("Fail - use wrong http method to send request", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/done?todoID=1", nil)
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(DoneHandler)
		handler.ServeHTTP(response, req)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusBadRequest, response.Code)
		}
	})
}

func TestUndoneHandler(t *testing.T) {
	t.Run("Success when update a resource to false", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/undone?todoID=1", nil)
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(UndoneHandler)
		handler.ServeHTTP(response, req)

		if status := response.Code; status != http.StatusOK {
			t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusOK, response.Code)
		}
	})

	t.Run("Fail - use wrong http method to send request", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/undone?todoID=1", nil)
		response := httptest.NewRecorder()
		handler := http.HandlerFunc(UndoneHandler)
		handler.ServeHTTP(response, req)

		if status := response.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusBadRequest, response.Code)
		}
	})
}
