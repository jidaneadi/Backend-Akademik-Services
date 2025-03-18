package routes

import (
	"net/http"
	"project-sia/controllers/authcontrollers"
)

func AuthRoutes(a authcontrollers.AuthControllers, r *http.ServeMux) {

	r.HandleFunc("/auth/register", a.Register)

}
