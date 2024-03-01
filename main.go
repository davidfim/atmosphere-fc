package main

import (
	"net/http"

	"github.com/davidfim/atmosphere-fc/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
