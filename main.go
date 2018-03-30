package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Some shit happen to server: %s", err.Error())
	}
}
