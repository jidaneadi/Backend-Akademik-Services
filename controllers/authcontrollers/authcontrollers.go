package authcontrollers

import (
	"net/http"
)

type AuthControllers interface {
	Register(w http.ResponseWriter, rq *http.Request)
}
