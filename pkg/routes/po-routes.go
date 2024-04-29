package routes

import (
	"net/http"

	"github.com/davlet61/powerapi-go/pkg/controllers"
)

var PoRoutes = func(r *http.ServeMux) {
	r.HandleFunc("/po/auth", controllers.AuthHandler)
}
