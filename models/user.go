package models

//easyjson:json
type User struct {
	Nickname string `json:"nickname,omitempty"`
	About    string `json:"about,omitempty"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}
