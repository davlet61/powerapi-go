package controllers

import (
	"fmt"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AuthHandler is working!")
}
