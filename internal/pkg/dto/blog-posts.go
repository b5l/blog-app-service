package dto

type BlogPostsResponseBody struct {
	Data []BlogPostsObject `json:"data"`
}

type BlogPostsObject struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type" db:"type"`
}
