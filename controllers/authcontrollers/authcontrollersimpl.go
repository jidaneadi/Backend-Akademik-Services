package authcontrollers

import (
	"net/http"
	"project-sia/exceptions"
	"project-sia/helpers"
	"project-sia/models/request"
	"project-sia/models/response"
	"project-sia/services/authservices"
)

type AuthControllersImpl struct {
	AuthServices authservices.AuthServices
}

func NewAuthControllers(a authservices.AuthServices) AuthControllers {
	return &AuthControllersImpl{
		AuthServices: a,
	}
}

func (s *AuthControllersImpl) Register(w http.ResponseWriter, rq *http.Request) {
	if rq.Method != http.MethodPost {
		http.Error(w, "Method NotAllowed", http.StatusMethodNotAllowed)
		return
	}

	if rq.Header.Get("Content-Type") != "application/json" {
		panic(exceptions.NewErrorUnsupported("Invalid Content-Type, expected application/json"))
	}
	reqBody := new(request.CreateSiswaNew)

	helpers.ReadReqBody(rq, reqBody)

	authServices := s.AuthServices.Register(rq.Context(), *reqBody)
	res := response.RegisterSucces{
		Meta: response.Meta{
			Code:    200,
			Status:  "SUCCES",
			Message: "Register sukses",
		},
		Data: authServices,
	}
	helpers.WriteToResBody(w, res)
}
