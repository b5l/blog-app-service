package model

type Login struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
