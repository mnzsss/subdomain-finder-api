package main

import (
	"net/http"
	handler "subdomain-finder-api/api"
)

func main() {
	http.HandleFunc("/", handler.Handler)

	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		panic(err)
	}
}
