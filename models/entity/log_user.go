package entity

import (
	"time"

	"github.com/google/uuid"
)

type Table_Log_User struct {
	ID         uuid.UUID `gorm:"primaryKey" json:"id"`
	Email      string    `gorm:"unique" json:"email"`
	Role       string    `gorm:"default='S'" json:"role"`
	Status     string    `gorm:"default='A'" json:"status"`
	No_Hp      uint64    `json:"no_hp"`
	Pass       string    `json:"pass"`
	Konf_Pass  string    `json:"konf_pass"`
	CreatedAt  time.Time `json:"created_at"`
	RiwayatLog time.Time `json:"riwayat_log"`
	Siswa      *Tb_Siswa `gorm:"foreignKey:Id_User;References:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"siswa"`
}

func (Table_Log_User) TableName() string {
	return "log_user"
}
