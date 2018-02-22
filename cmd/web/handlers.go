package web

import (
	"net/http"

	"github.com/lccezinha/twodo/cmd/environment"
)

var app *environment.Application

func init() {
	app = environment.Init()
}

// IndexHandler will handle request to "/" path
func IndexHandler(w http.ResponseWriter, r *http.Request) {

}

// CreateHandler will handle request to "/" path with post request
func CreateHandler(w http.ResponseWriter, r *http.Request) {

}

// DestroyHandler will destroy a single resource
func DestroyHandler(w http.ResponseWriter, r *http.Request) {

}

// DoneHandler will make resource as done
func DoneHandler(w http.ResponseWriter, r *http.Request) {

}

// UndoneHandler will make resource as done
func UndoneHandler(w http.ResponseWriter, r *http.Request) {

}
