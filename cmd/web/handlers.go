package web

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

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

	todos, err := app.ListService.Run()

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	t.Execute(w, todos)
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

	http.Redirect(w, r, "/", http.StatusFound)
}

// DestroyHandler will destroy a single resource
func DestroyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.Header().Set("Allow", "DELETE")
		w.Header().Set("error_message", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	todoParamID := r.URL.Query().Get("todoID")
	todoID, _ := strconv.Atoi(todoParamID)

	err := app.DestroyService.Run(todoID)

	if err != nil {
		w.Header().Set("error_message", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusNoContent)
}

// DoneHandler will make resource as done
func DoneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.Header().Set("error_message", "Wrong http method in request")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	todoParamID := r.URL.Query().Get("todoID")
	todoID, _ := strconv.Atoi(todoParamID)

	_, err := app.UpdateService.Run(todoID, true)

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err.Error())
	}

	http.Redirect(w, r, "/", http.StatusOK)
}

// UndoneHandler will make resource as done
func UndoneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.Header().Set("error_message", "Wrong http method in request")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	todoParamID := r.URL.Query().Get("todoID")
	todoID, _ := strconv.Atoi(todoParamID)

	_, err := app.UpdateService.Run(todoID, false)

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err.Error())
	}

	http.Redirect(w, r, "/", http.StatusOK)
}
