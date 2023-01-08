package dto

type BlogPostsResponseBody struct {
	Data []BlogPostsObject `json:"data"`
}

type BlogPostsObject struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	ShortDescription string `json:"shortDescription" db:"short_description"`
}
