package req

type ReqUserUpdate struct {
	FullName string `json:"full_name" validate:"required"`
	Email string `json:"email" validate:"email"`
}
