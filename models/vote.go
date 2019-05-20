package models

//easyjson:json
type Vote struct {
	Voice    int    `json:"voice"`
	Nickname string `json:"nickname"`
	Id       int
}
