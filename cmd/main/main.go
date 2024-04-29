package main

import (
	"log"
	"net/http"

	"github.com/davlet61/powerapi-go/pkg/routes"
)

func main() {
	r := http.NewServeMux()
	routes.PoRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
