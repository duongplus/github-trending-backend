package model

import "time"

type User struct {
	UserId    string    `db:"user_id, omitempty" json:"-"`
	FullName  string    `db:"full_name, omitempty" json:"full_name"`
	Email     string    `db:"email, omitempty" json:"email"`
	Password  string    `db:"password, omitempty" json:"-"`
	Role      string    `db:"role, omitempty" json:"role"`
	CreatedAt time.Time `db:"created_at, omitempty" json:"-"`
	UpdatedAt time.Time `db:"updated_at, omitempty" json:"-"`
	Token     string    `json:"token,omitempty"`
}
