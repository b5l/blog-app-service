package model

type BlogPosts struct {
	Id               string `db:"id"`
	Title            string `db:"title"`
	ShortDescription string `db:"short_description"`
}
