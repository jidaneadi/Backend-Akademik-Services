package services

import (
	"backend-sia/models/request"
	"backend-sia/models/response"
	"context"
)

type AuthServices interface {
	Register(c context.Context, r request.CreateSiswaNew) response.DataRegisterSucces
}
