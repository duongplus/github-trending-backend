package req

type ReqSignIn struct {
	Email    string `validate:"email" json:"email"`
	Password string `validate:"required,pwd" json:"password"`
}
