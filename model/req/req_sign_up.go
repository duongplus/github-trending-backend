package req

type ReqSignUp struct {
	FullName string `validate:"required" json:"full_name"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"pwd" json:"password"`
}
