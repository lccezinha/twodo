package main

import (
	"net/http"

	"github.com/lccezinha/twodo/cmd/web"
)

func main() {
	http.HandleFunc("/", web.IndexHandler)
	http.ListenAndServe(":8080", nil)
}
