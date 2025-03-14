package services

import (
	"backend-sia/databases"
	"backend-sia/helpers"
	"backend-sia/models/entity"
	"backend-sia/models/request"
	"backend-sia/models/response"
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthServicesImpl struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func (auth *AuthServicesImpl) Register(c context.Context, r request.CreateSiswaNew) response.DataRegisterSucces {
	validation := auth.Validate.Struct(r)
	helpers.PanicErr(validation)

	if r.Email != "" {
		helpers.ErrorBadRequest(errors.New("email tidak boleh kosong"))
	}
	if r.Pass != "" {
		helpers.ErrorBadRequest(errors.New("password tidak boleh kosong"))
	}
	if r.KonfPass != "" {
		helpers.ErrorBadRequest(errors.New("konfirmasi password tidak boleh kosong"))
	}
	if r.NoHp <= 0 {
		helpers.ErrorBadRequest(errors.New("nomor hp tidak boleh kosong"))
	}
	if r.Nisn <= 0 {
		helpers.ErrorBadRequest(errors.New("nisn tidak boleh kosong"))
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
	if r.Pass != r.KonfPass {
		helpers.ErrorBadRequest(errors.New("password dan konfirmasi password harus sama"))
	}

	createUser := entity.Table_Log_User{
		ID:        uuid.New(),
		Email:     r.Email,
		Status:    "A",
		No_Hp:     r.NoHp,
		Pass:      r.Pass,
		Konf_Pass: r.KonfPass,
		CreatedAt: time.Now(),
	}
	createSiswa := entity.Tb_Siswa{
		Nisn:         r.Nisn,
		Id_User:      createUser.ID,
		Nama:         r.Nama,
		Tempat_Lahir: r.Tempat_Lahir,
		Tgl_Lahir:    r.Tgl_Lahir,
		Alamat:       r.Alamat,
		Jns_Kelamin:  r.Jns_Kelamin,
		Agama:        r.Agama,
		Ket_Lulus:    r.Ket_Lulus,
		Created_At:   time.Now(),
		Updated_At:   time.Now(),
	}
	cU := databases.DB.Create(createUser).Error
	helpers.ErrorBadRequest(cU)

	cS := databases.DB.Create(createSiswa).Error
	helpers.ErrorBadRequest(cS)

	return helpers.ToRegisterResponse(r)
}
