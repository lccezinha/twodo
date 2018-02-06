package web

import (
	"errors"
	"html/template"
	"log"
	"net/http"
)

var errTemplateNotFound = errors.New("Template not found")

// IndexHandler will handle request to "/" path
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../../views/index.html")

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		// panic(errTemplateNotFound)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	t.Execute(w, nil)
}

// CreateHandler will handle request to "/" path with post request
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//
	// if r.Method != "POST" {
	// 	w.Header().Set("Allow", "POST")
	// 	w.Header().Set("Error Message", "Use wrong http method")
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// }
	//
	// title := r.FormValue("title")
	// description := r.FormValue("description")
	//
	// fmt.Println("Values", title, description)
}
