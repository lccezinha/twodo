package web

import (
	"html/template"
	"log"
	"net/http"

	"github.com/lccezinha/twodo/cmd/environment"
)

var app *environment.Application

func init() {
	app = environment.Init()
}

// IndexHandler will handle request to "/" path
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/home/luizcezer/go/src/github.com/lccezinha/twodo/views/index.html")

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	t.Execute(w, nil)
}

// CreateHandler will handle request to "/" path with post request
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.Header().Set("error_message", "Wrong http method in request")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	description := r.FormValue("description")

	err := app.CreateService.Run(title, description)

	if err != nil {
		w.Header().Set("error_message", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// func DestroyHandler(w http.ResponseWriter, r *http.Request) {}
// func UpdateHandler(w http.ResponseWriter, r *http.Request) {}
