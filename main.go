package main

import (
	"net/http"

	"github.com/flaviofilipe/go_login_encoders/routes"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8000", nil)
}
