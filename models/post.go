package models

import "time"

type Post struct {
	Author    string    `json:"author"`
	Created   time.Time `json:"created"`
	Forum     string    `json:"forum, omitempty"`
	Id        int       `json:"id, omitempty"`
	IsEdited  bool      `json:"isEdited, omitempty"`
	Message   string    `json:"message"`
	Parent    int       `json:"parent, omitempty"`
	ThreadId  int       `json:"thread, omitempty"`
	Thread    string
	Path      string
	Childrens int
}

// ключ = types.Post.Id
type PostConnectionsItem struct {
	ThreadId         int    // types.Post.ThreadId
	MaterializedPath string // types.Post.MaterializedPath
	NumberOfChildren int    // types.Post.NumberOfChildren
}
