package request

type LoginUser struct {
	Email string `validate:"email,min=11,max=30" json:"email"`
	Pass  string `json:"password"`
}
type ReqToken struct {
	AccesToken   string `json:"acces_token"`
	RefreshToken string `json:"refresh_token"`
}
