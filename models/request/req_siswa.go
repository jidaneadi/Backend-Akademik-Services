package request

import "time"

type CreateSiswaNew struct {
	Nisn         uint64    `validate:"min=16,max=19" json:"nisn"`
	Id_User      string    `validate:"min=16,max=40" json:"id_user"`
	Nama         string    `validate:"min=3,max=30" json:"nama"`
	Tempat_Lahir string    `validate:"min=3,max=15" json:"tempat_lahir"`
	Tgl_Lahir    time.Time `validate:"required" json:"tgl_lahir"`
	Alamat       string    `validate:"min=16,max=200" json:"alamat"`
	Jns_Kelamin  string    `validate:"max=1" json:"jns_kelamin"`
	Agama        string    `validate:"max=1" json:"agama"`
	Ket_Lulus    string    `validate:"max=1" json:"ket_lulus"`
}
