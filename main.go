package main

import (
	"log"
	"net/http"

	"github.com/lccezinha/twodo/cmd/web"
)

func main() {
	http.HandleFunc("/", web.IndexHandler)
	http.HandleFunc("/todos", web.CreateHandler)
	http.HandleFunc("/destroy", web.DestroyHandler)
	http.HandleFunc("/done", web.DoneHandler)
	http.HandleFunc("/undone", web.UndoneHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Print("Running server... http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Some shit happen to server: %s", err.Error())
	}
}
