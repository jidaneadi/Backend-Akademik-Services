package authservices

import (
	"context"
	"project-sia/models/request"
	"project-sia/models/response"
)

type AuthServices interface {
	Register(c context.Context, rq request.CreateSiswaNew) response.DataRegisterSucces
	Login(c context.Context, rq request.LoginUser) response.DataLoginSucces
}
