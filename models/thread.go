package models

import "time"

type Thread struct {
	Author   string    `json:"author"`
	Created  time.Time `json:"created"`
	Forum    string    `json:"forum,omitempty"`
	Slug     string    `json:"slug,omitempty"`
	Id       int       `json:"id,omitempty"`
	Message  string    `json:"message"`
	Title    string    `json:"title,omitempty"`
	Votes    int       `json:"votes,omitempty"`
	Children int       `json:"-"`
}
