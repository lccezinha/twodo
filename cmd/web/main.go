package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lccezinha/twodo/cmd/web/application"
)

func main() {
	router := application.InitializeRouter()
	port := ":8888"

	fmt.Printf("Server running on port: %s\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
