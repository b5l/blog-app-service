package dto

type BlogDetailsResponseBody struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Type        string `json:"type" db:"type"`
	Description string `json:"description" db:"description"`
}
