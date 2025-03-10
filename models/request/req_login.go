package request

type RequestLogin struct {
	Email    string `validate:"email, min=10"`
	Password string `validate:"min=8"`
}
