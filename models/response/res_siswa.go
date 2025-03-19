package response

import (
	"time"

	"github.com/google/uuid"
)

type DataSiswaWLog struct {
	ID           uuid.UUID `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"unique" json:"email"`
	Role         string    `gorm:"default='S'" json:"role"`
	Status       string    `gorm:"default='A'" json:"status"`
	No_Hp        uint64    `json:"no_hp"`
	Pass         string    `json:"pass"`
	Konf_Pass    string    `json:"konf_pass"`
	Nisn         uint64    `gorm:"primaryKey" json:"nisn"`
	Nama         string    `json:"nama"`
	Tempat_Lahir string    `json:"tempat_lahir"`
	Tgl_Lahir    time.Time `json:"tgl_lahir"`
	Alamat       string    `json:"alamat"`
	Jns_Kelamin  string    `gorm:"default='L'" json:"jns_kelamin"`
	Agama        string    `gorm:"default='I'" json:"agama"`
	Ket_Lulus    string    `gorm:"default='B'" json:"ket_lulus"`
	CreatedAt    time.Time `json:"created_at"`
	RiwayatLog   time.Time `json:"riwayat_log"`
}

func (d DataSiswaWLog) TableName() string {
	return "log_user"
}
