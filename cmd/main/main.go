package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davlet61/powerapi-go/pkg/routes"
)

func main() {
	r := http.NewServeMux()
	routes.PoRoutes(r)
	http.Handle("/", r)
	port := 8080
	log.Println("Starting server on port " + fmt.Sprint(port) + "\r")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
