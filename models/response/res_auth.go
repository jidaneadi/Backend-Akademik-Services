package response

import (
	"time"
)

type Token struct {
	AccesJwt string `json:"acces_jwt"`
	Refresh  string `json:"refresh_jwt"`
}

type RegisterSucces struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type DataRegisterSucces struct {
	Email        string    `json:"email"`
	No_Hp        uint64    `json:"no_hp"`
	Status       string    `json:"status"`
	Nisn         uint64    `json:"nisn"`
	Id_User      string    `json:"id_user"`
	Nama         string    `json:"nama"`
	Tempat_Lahir string    `json:"tempat_lahir"`
	Tgl_Lahir    time.Time `json:"tgl_lahir"`
	Alamat       string    `json:"alamat"`
	Jns_Kelamin  string    `json:"jns_kelamin"`
	Agama        string    `json:"agama"`
	Ket_Lulus    string    `json:"ket_lulus"`
	Created_At   time.Time `json:"created_at"`
	Updated_At   time.Time `json:"updated_at"`
}

type LoginSucces struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}
