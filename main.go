package main

import (
	"net/http"

	"github.com/web-app-go/infra/routes"
)

func main() {

	routes.HandleRequests()
	http.ListenAndServe(":8080", nil)


}
