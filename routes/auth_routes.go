package routes

import (
	"backend-sia/controllers"
	"backend-sia/middlewares"
	"net/http"
)

func AuthRoutes(r http.ServeMux) {
	r.HandleFunc("/login", middlewares.WithErrorHandler(controllers.LoginControllers))
}
