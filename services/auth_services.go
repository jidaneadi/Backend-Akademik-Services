package services

import (
	"backend-sia/helpers"
	"backend-sia/models/request"
	"backend-sia/models/response"
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthServicesImpl struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func (auth *AuthServicesImpl) Register(c context.Context, r request.CreateSiswaNew) response.DataRegisterSucces {
	validation := auth.Validate.Struct(r)
	helpers.PanicErr(validation)

	if r.Nisn <= 0 {
		helpers.ErrorBadRequest(errors.New("nisn tidak boleh kosong"))
	}
	if r.Id_User != "" {
		helpers.ErrorBadRequest(errors.New("id user tidak boleh kosong"))
	}
	if r.Nama != "" {
		helpers.ErrorBadRequest(errors.New("nama tidak boleh kosong"))
	}
	if r.Tempat_Lahir == "" {
		helpers.ErrorBadRequest(errors.New("tempat lahir tidak boleh kosong"))
	}
	if r.Alamat == "" {
		helpers.ErrorBadRequest(errors.New("alamat tidak boleh kosong"))
	}
	if r.Jns_Kelamin == "" {
		helpers.ErrorBadRequest(errors.New("jenis kelamin tidak boleh kosong"))
	}
	if r.Agama == "" {
		helpers.ErrorBadRequest(errors.New("agama tidak boleh kosong"))
	}
	if r.Ket_Lulus == "" {
		helpers.ErrorBadRequest(errors.New("keterangan lulus tidak boleh kosong"))
	}

	return helpers.ToRegisterResponse(r)
}
