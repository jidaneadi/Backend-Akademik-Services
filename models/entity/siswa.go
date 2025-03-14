package entity

import (
	"time"

	"github.com/google/uuid"
)

type Tb_Siswa struct {
	Nisn         uint64    `gorm:"primaryKey,allowNull=false" json:"nisn"`
	Id_User      uuid.UUID `json:"id_user"`
	Nama         string    `json:"nama"`
	Tempat_Lahir string    `json:"tempat_lahir"`
	Tgl_Lahir    time.Time `json:"tgl_lahir"`
	Alamat       string    `json:"alamat"`
	Jns_Kelamin  string    `gorm:"default=L" json:"jns_kelamin"`
	Agama        string    `gorm:"default=I" json:"agama"`
	Ket_Lulus    string    `gorm:"default=B" json:"ket_lulus"`
	Created_At   time.Time `json:"created_at"`
	Updated_At   time.Time `json:"updated_at"`
}

func (Tb_Siswa) TableName() string {
	return "siswa"
}
