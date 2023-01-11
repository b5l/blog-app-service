package model

type BlogCreate struct {
	Title            string `db:"title"`
	ShortDescription string `db:"short_description"`
	LongDescription  string `db:"long_description"`
}
