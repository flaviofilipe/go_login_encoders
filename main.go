package main

import (
	"net/http"

	"github.com/flaviofilipe/login/routes"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8000", nil)
}
