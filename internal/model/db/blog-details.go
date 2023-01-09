package model

type BlogDetails struct {
	Id               string `db:"id"`
	Title            string `db:"title"`
	ShortDescription string `db:"short_description"`
	LongDescription  string `db:"long_description"`
}
