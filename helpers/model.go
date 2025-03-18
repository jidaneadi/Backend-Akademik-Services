package helpers

import (
	"project-sia/models/request"
	"project-sia/models/response"
	"time"
)

func ToRegisterResponse(r request.CreateSiswaNew) response.DataRegisterSucces {
	parseTgl, err := time.Parse("2006-01-02", r.Tgl_Lahir)
	if err != nil {
		panic(err)
	}
	return response.DataRegisterSucces{
		Email:        r.Email,
		No_Hp:        r.NoHp,
		Nisn:         r.Nisn,
		Status:       "A",
		Nama:         r.Nama,
		Tempat_Lahir: r.Tempat_Lahir,
		Tgl_Lahir:    parseTgl,
		Alamat:       r.Alamat,
		Jns_Kelamin:  r.Jns_Kelamin,
		Agama:        r.Agama,
		Ket_Lulus:    r.Ket_Lulus,
		Created_At:   time.Now(),
		Updated_At:   time.Now(),
	}
}
