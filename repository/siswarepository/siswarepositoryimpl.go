package siswarepository

import (
	"project-sia/databases"
	"project-sia/exceptions"
	"project-sia/helpers"
	"project-sia/models/request"
	"project-sia/models/response"

	"gorm.io/gorm"
)

type SiswaRepositoryImpl struct {
}

func NewSiswaRepository() SiswaRepositoryImpl {
	return SiswaRepositoryImpl{}
}

func (r SiswaRepositoryImpl) GetSiswaByEmail(rq request.GetByData, data *response.DataSiswaWLog) *response.DataSiswaWLog {
	if err := databases.DB.Where("log_user.email = ?", rq.Email).
		Select(`log_user.id AS id, log_user.email, log_user.pass, 
	         siswa.nisn, siswa.nama, siswa.tempat_lahir, siswa.tgl_lahir, 
	         siswa.alamat, siswa.jns_kelamin, siswa.agama, siswa.ket_lulus, 
	         log_user.role, log_user.status, log_user.no_hp, 
	         log_user.konf_pass, log_user.created_at, log_user.riwayat_log`).
		Joins("LEFT JOIN siswa ON siswa.id_user = log_user.id").
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(exceptions.NewNotFoundError("email atau password salah"))
		}
		helpers.PanicErr(err)
	}
	return data
}
