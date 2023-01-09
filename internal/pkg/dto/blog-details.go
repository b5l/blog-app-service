package dto

type BlogDetailsResponseBody struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	ShortDescription string `json:"shortDescription" db:"short_description"`
	LongDescription  string `json:"longDescription" db:"long_description"`
}
