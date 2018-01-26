package web

import (
	"errors"
	"html/template"
	"net/http"
)

var errTemplateNotFound = errors.New("Template not found")

// IndexHandler will handle request to "/" path
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("../../views/index.html")

	if err != nil {
		panic(errTemplateNotFound)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	template.Execute(w, nil)
}

// CreateHandler will handle request to "/" path with post request
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.Header().Set("Error Message", "Use wrong http method")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
