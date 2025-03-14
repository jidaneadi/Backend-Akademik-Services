package helpers

import (
	"backend-sia/models/request"
	"backend-sia/models/response"
	"time"
)

func ToRegisterResponse(r request.CreateSiswaNew) response.DataRegisterSucces {
	return response.DataRegisterSucces{
		Email:        r.Email,
		No_Hp:        r.NoHp,
		Nisn:         r.Nisn,
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
}
