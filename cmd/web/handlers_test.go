package web

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
	// t.Run("Fail - use wrong http method to send request", func(t *testing.T) {
	// 	data := url.Values{}
	// 	data.Add("title", "")
	// 	data.Add("description", "Some Description")
	// 	req := httptest.NewRequest("GET", "/create", strings.NewReader(data.Encode()))
	// 	w := httptest.NewRecorder()
	// 	handler := http.HandlerFunc(CreateHandler)
	// 	handler.ServeHTTP(w, req)
	//
	// 	if status := w.Code; status != http.StatusMethodNotAllowed {
	// 		t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusBadRequest, w.Code)
	// 	}
	// })

	// t.Run("Fail because todo title is not set", func(t *testing.T) {
	// 	params := url.Values{}
	// 	params.Add("title", "")
	// 	params.Add("description", "Some Description")
	// 	req, _ := http.NewRequest("POST", "/", strings.NewReader(params.Encode()))
	// 	w := httptest.NewRecorder()
	// 	handler := http.HandlerFunc(CreateHandler)
	// 	handler.ServeHTTP(w, req)
	//
	// 	if status := w.Code; status != http.StatusBadRequest {
	// 		t.Errorf("Handler returning wrong http status code, expected %v, received %v", http.StatusBadRequest, w.Code)
	// 	}
	// })
}
