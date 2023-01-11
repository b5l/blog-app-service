package model

type BlogDetails struct {
	Id               int    `db:"id"`
	Title            string `db:"title"`
	ShortDescription string `db:"short_description"`
	LongDescription  string `db:"long_description"`
}
