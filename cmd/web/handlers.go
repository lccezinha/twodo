package web

import (
	"errors"
	"html/template"
	"net/http"
)

var errTemplateNotFound = errors.New("Template not found")

// IndexHandler will handle request to "/" path
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template, err := template.ParseFiles("../../views/index.html")

	if err != nil {
		panic(errTemplateNotFound)
	}

	template.Execute(w, nil)
}
