package authcontrollers

import (
	"backend-sia/helpers"
	"backend-sia/models/request"
	"backend-sia/models/response"
	"backend-sia/services"
	"net/http"
)

type AuthControllersImpl struct {
	AuthServices services.AuthServices
}

func (auth *AuthControllersImpl) Register(w http.ResponseWriter, r *http.Request) {
	rB := new(request.CreateSiswaNew)
	helpers.ReadFromRequestBody(r, rB)

	authResponse := auth.AuthServices.Register(r.Context(), *rB)
	webResponse := response.RegisterSucces{
		Meta: response.Meta{
			Code:    200,
			Status:  false,
			Message: "Regristasi sukses!",
		},
		Data: authResponse,
	}

	helpers.WriteToResponseBody(w, webResponse)
}
