package controllers

import (
	"backend-sia/database"
	"backend-sia/helpers"
	"backend-sia/models/entity"
	"backend-sia/models/request"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginControllers(res http.ResponseWriter, req *http.Request) error {
	rB := new(request.RequestLogin)
	if rB.Email == "" {
		return helpers.MessageError(res, 400, "Email tidak boleh kosong", false)
	}

	if rB.Password == "" {
		return helpers.MessageError(res, 400, "Password tidak boleh kosong", false)
	}

	if len(rB.Password) <= 8 {
		return helpers.MessageError(res, 400, "Password harus memiliki minimal 8 karakter", false)
	}

	cD := new(entity.Table_Log_User)
	if err := database.DB.Where("email =?", rB.Email).First(cD).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.MessageError(res, 404, "Email or password invalid", false)
		}
		return helpers.MessageError(res, 500, err.Error(), false)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(cD.Pass), []byte(rB.Password)); err != nil {
		return helpers.MessageError(res, 400, "Email or pasword invalid", false)
	}

	return helpers.MessageSucces(res, "Login succes....", true)
}
