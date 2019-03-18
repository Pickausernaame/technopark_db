package models

type Post struct {
	Author   string `json:"author"`
	Created  string `json:"created"`
	Forum    string `json:"forum, omitempty"`
	Id       int    `json:"id, omitempty"`
	IsEdited bool   `json:"isEdited, omitempty"`
	Message  string `json:"message"`
	Parent   int    `json:"parent, omitempty"`
	ThreadId int    `json:"thread, omitempty"`
}
