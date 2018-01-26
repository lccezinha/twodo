package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)
	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusOK, w.Code)
	}
}

func TestCreateHandler(t *testing.T) {
	t.Run("Fail - use wrong http method to send request", func(t *testing.T) {
		data := url.Values{}
		data.Add("title", "")
		data.Add("description", "Some Description")
		req := httptest.NewRequest("GET", "/", strings.NewReader(data.Encode()))
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateHandler)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusBadRequest, w.Code)
		}
	})

	// t.Run("Fail because todo title is not set", func(t *testing.T) {
	// 	data := url.Values{}
	// 	data.Add("title", "")
	// 	data.Add("description", "Some Description")
	// 	req := httptest.NewRequest("POST", "/", strings.NewReader(data.Encode()))
	// 	w := httptest.NewRecorder()
	// 	handler := http.HandlerFunc(CreateHandler)
	// 	handler.ServeHTTP(w, req)
	//
	// 	if status := w.Code; status != http.StatusBadRequest {
	// 		t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusBadRequest, w.Code)
	// 	}
	// })
}
