package model

type BlogCreate struct {
	Title       string `db:"title"`
	Type        string `db:"type"`
	Description string `db:"description"`
}
