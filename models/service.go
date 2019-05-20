package models

//easyjson:json
type Service struct {
	Users   int `json:"user"`
	Forums  int `json:"forum"`
	Posts   int `json:"post"`
	Threads int `json:"thread"`
}
