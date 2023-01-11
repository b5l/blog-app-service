package model

type BlogDetails struct {
	Id          int    `db:"id"`
	Title       string `db:"title"`
	Type        string `db:"type"`
	Description string `db:"description"`
}
