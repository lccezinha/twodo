package main

import (
	"log"
	"net/http"

	"github.com/lccezinha/twodo/cmd/web/application"
)

func main() {
	router := application.InitializeRouter()
	port := ":8888"

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
	log.Printf("Server running in port: %s\n", port)
}
