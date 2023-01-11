package model

type BlogPosts struct {
	Id    string `db:"id"`
	Title string `db:"title"`
	Type  string `db:"type"`
}
