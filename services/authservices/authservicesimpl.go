package authservices

import (
	"context"
	"project-sia/databases"
	"project-sia/helpers"
	"project-sia/models/entity"
	"project-sia/models/request"
	"project-sia/models/response"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthServicesImpl struct {
	Validate *validator.Validate
}

func NewAuthServices(v *validator.Validate) AuthServices {
	return &AuthServicesImpl{
		Validate: v,
	}
}

func (s *AuthServicesImpl) Register(c context.Context, rq request.CreateSiswaNew) response.DataRegisterSucces {
	v := s.Validate.Struct(rq)
	helpers.PanicErr(v)

	hashPass, err := bcrypt.GenerateFromPassword([]byte(rq.Pass), 10)
	if err != nil {
		helpers.PanicErr(err)
	}
	hashKonfPass, err := bcrypt.GenerateFromPassword([]byte(rq.KonfPass), 10)
	if err != nil {
		helpers.PanicErr(err)
	}
	createUser := entity.Table_Log_User{
		ID:         uuid.New(),
		Email:      rq.Email,
		Status:     "A",
		No_Hp:      rq.NoHp,
		Pass:       string(hashPass),
		Konf_Pass:  string(hashKonfPass),
		CreatedAt:  time.Now(),
		RiwayatLog: time.Now(),
	}
	parseTgl, err := time.Parse("2006-01-02", rq.Tgl_Lahir)
	if err != nil {
		panic(err)
	}

	createSiswa := entity.Tb_Siswa{
		Nisn:         rq.Nisn,
		Id_User:      createUser.ID,
		Nama:         rq.Nama,
		Tempat_Lahir: rq.Tempat_Lahir,
		Tgl_Lahir:    parseTgl,
		Alamat:       rq.Alamat,
		Jns_Kelamin:  rq.Jns_Kelamin,
		Agama:        rq.Agama,
		Ket_Lulus:    rq.Ket_Lulus,
		Created_At:   time.Now(),
		Updated_At:   time.Now(),
	}
	if err := databases.DB.Create(createUser).Error; err != nil {
		helpers.PanicErr(err)
	}

	if err := databases.DB.Create(createSiswa).Error; err != nil {
		helpers.PanicErr(err)
	}

	return helpers.ToRegisterResponse(rq)
}
