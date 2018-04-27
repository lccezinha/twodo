package main

import (
	"log"
	"net/http"

	"github.com/lccezinha/twodo/cmd/web/app"
)

func main() {
	router := app.InitializeRouter()
	port := ":8080"

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
	log.Printf("Server running in port: %s\n", port)
}
