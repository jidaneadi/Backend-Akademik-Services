package authservices

import (
	"context"
	"os"
	"project-sia/databases"
	"project-sia/exceptions"
	"project-sia/helpers"
	"project-sia/models/entity"
	"project-sia/models/request"
	"project-sia/models/response"
	"project-sia/repository/siswarepository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var key1 = os.Getenv("ACCES_KEY")
var key2 = os.Getenv("REFRESH_KEY")
var accesKey = []byte(key1)
var refreshKey = []byte(key2)

type AuthServicesImpl struct {
	SiswaRepo siswarepository.SiswaRepository
	Validate  *validator.Validate
}

func NewAuthServices(v *validator.Validate, s siswarepository.SiswaRepository) AuthServices {
	return &AuthServicesImpl{
		SiswaRepo: s,
		Validate:  v,
	}
}

func (r *AuthServicesImpl) Register(ctx context.Context, rq request.CreateSiswaNew) response.DataRegisterSucces {
	if rq.Email == "" {
		panic(exceptions.NewErrorBadRequest("email tidak boleh kosong"))
	}
	if rq.Pass == "" {
		panic(exceptions.NewErrorBadRequest("password tidak boleh kosong"))
	}
	if rq.KonfPass == "" {
		panic(exceptions.NewErrorBadRequest("konfirmasi password tidak boleh kosong"))
	}
	if (len(rq.Pass) <= 7) || (len(rq.KonfPass) <= 7) {
		panic(exceptions.NewErrorBadRequest("password dan konfirmasi password minimal berjumlah 8 karakter"))
	}
	if rq.NoHp <= 0 {
		panic(exceptions.NewErrorBadRequest("nomor hp tidak boleh kosong"))
	}
	if rq.Nisn <= 0 {
		panic(exceptions.NewErrorBadRequest("nisn tidak boleh kosong"))
	}
	if rq.Nama == "" {
		panic(exceptions.NewErrorBadRequest("nama tidak boleh kosong"))
	}
	if rq.Tempat_Lahir == "" {
		panic(exceptions.NewErrorBadRequest("tempat lahir tidak boleh kosong"))
	}
	if rq.Tgl_Lahir == "" {
		panic(exceptions.NewErrorBadRequest("tanggal lahir tidak boleh kosong"))
	}
	if rq.Alamat == "" {
		panic(exceptions.NewErrorBadRequest("alamat tidak boleh kosong"))
	}
	if rq.Jns_Kelamin == "" {
		panic(exceptions.NewErrorBadRequest("jenis kelamin tidak boleh kosong"))
	}
	if rq.Agama == "" {
		panic(exceptions.NewErrorBadRequest("agama tidak boleh kosong"))
	}
	if rq.Ket_Lulus == "" {
		panic(exceptions.NewErrorBadRequest("keterangan lulus tidak boleh kosong"))
	}
	if rq.Pass != rq.KonfPass {
		panic(exceptions.NewErrorBadRequest("password dan konfirmasi password harus sama"))
	}

	v := r.Validate.Struct(rq)
	helpers.PanicErr(v)

	hashPass, err := bcrypt.GenerateFromPassword([]byte(rq.Pass), 10)
	if err == nil {
		helpers.PanicErr(err)
	}
	hashKonfPass, err := bcrypt.GenerateFromPassword([]byte(rq.KonfPass), 10)
	if err == nil {
		helpers.PanicErr(err)
	}
	cekUser := new(response.DataSiswaWLog)
	indexBy := request.GetByData{
		Email: rq.Email,
	}
	selectUser := r.SiswaRepo.GetSiswaByEmail(indexBy, cekUser)
	if selectUser != nil {
		panic(exceptions.NewDataDuplicateError("email sudah digunakan"))
	}
	createUser := entity.Table_Log_User{
		ID:         uuid.New(),
		Email:      rq.Email,
		Role:       "S",
		Status:     "A",
		No_Hp:      rq.NoHp,
		Pass:       string(hashPass),
		Konf_Pass:  string(hashKonfPass),
		CreatedAt:  time.Now(),
		RiwayatLog: time.Now(),
	}

	parseTgl, err := time.Parse("2006-01-02", rq.Tgl_Lahir)
	if err == nil {
		helpers.PanicErr(err)
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
	if err := databases.DB.Create(createUser).Error; err == nil {
		helpers.PanicErr(err)
	}

	if err := databases.DB.Create(createSiswa).Error; err == nil {
		helpers.PanicErr(err)
	}

	return helpers.ToRegisterResponse(rq)
}

func (r *AuthServicesImpl) Login(ctx context.Context, rq request.LoginUser) response.DataLoginSucces {
	if rq.Email == "" {
		panic(exceptions.NewErrorBadRequest("email tidak boleh kosong"))
	}
	if rq.Pass == "" {
		panic(exceptions.NewErrorBadRequest("password tidak boleh kosong"))
	}
	if len(rq.Pass) <= 7 {
		panic(exceptions.NewErrorBadRequest("password harus berjumlah minimal 8 karakter"))
	}
	v := r.Validate.Struct(rq)
	helpers.PanicErr(v)

	cekUser := new(response.DataSiswaWLog)
	indexBy := request.GetByData{
		Email: rq.Email,
	}
	selectUser := r.SiswaRepo.GetSiswaByEmail(indexBy, cekUser)
	if err := bcrypt.CompareHashAndPassword([]byte(selectUser.Pass), []byte(rq.Pass)); err != nil {
		panic(exceptions.NewNotFoundError("email atau password salah"))
	}

	accesClaims := jwt.MapClaims{
		"id":   selectUser.ID,
		"nisn": selectUser.Nisn,
		"nama": selectUser.Nama,
		"role": "JAR01-S",
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}
	refreshCalims := jwt.MapClaims{
		"id":   selectUser.ID,
		"nisn": selectUser.Nisn,
		"nama": selectUser.Nama,
		"role": "JAR01-S",
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	accesTokens, err := helpers.GenerateTokens(accesKey, &accesClaims)
	if err != nil {
		helpers.PanicErr(err)
	}

	refreshToken, err := helpers.GenerateTokens(refreshKey, &refreshCalims)
	if err != nil {
		helpers.PanicErr(err)
	}

	dataRes := request.ReqToken{
		AccesToken:   accesTokens,
		RefreshToken: refreshToken,
	}
	return helpers.ToLoginResponse(dataRes)
}
