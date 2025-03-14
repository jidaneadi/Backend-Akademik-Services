package request

import "time"

type CreateSiswaNew struct {
	Email        string    `validate:"email,min=11,max=30" json:"email"`
	NoHp         uint64    `validate:"min=6000000000,max=180000000000" json:"no_hp"`
	Pass         string    `validate:"min=8,max=20" json:"password"`
	KonfPass     string    `validate:"min=8,max=20" json:"konf_password"`
	Nisn         uint64    `validate:"min=16,max=19" json:"nisn"`
	Nama         string    `validate:"min=3,max=30" json:"nama"`
	Tempat_Lahir string    `validate:"min=3,max=15" json:"tempat_lahir"`
	Tgl_Lahir    time.Time `validate:"required" json:"tgl_lahir"`
	Alamat       string    `validate:"min=16,max=200" json:"alamat"`
	Jns_Kelamin  string    `validate:"max=1" json:"jns_kelamin"`
	Agama        string    `validate:"max=1" json:"agama"`
	Ket_Lulus    string    `validate:"max=1" json:"ket_lulus"`
}
