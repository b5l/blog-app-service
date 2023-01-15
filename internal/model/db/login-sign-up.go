package model

type LoginSignUp struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
