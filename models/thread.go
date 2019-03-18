package models

type Thread struct {
	Author  string `json:"author"`
	Created string `json:"created"`
	Forum   string `json:"forum,omitempty"`
	Id      string `json:"id,omitempty"`
	Message string `json:"message"`
	Title   string `json:"title,omitempty"`
	Votes   string `json:"votes,omitempty"`
}
