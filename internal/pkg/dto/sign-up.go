package dto

type SignUpResponseBody struct {
	UserTaken    bool `json:"userTaken"`
	IsSuccessful bool `json:"isSuccessful"`
}
