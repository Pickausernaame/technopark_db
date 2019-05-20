package models

import "time"

//easyjson:json
type Posts []Post

//easyjson:json
type Post struct {
	Author    string    `json:"author"`
	Created   time.Time `json:"created,omitempty"`
	Forum     string    `json:"forum,omitempty"`
	Id        int       `json:"id,omitempty"`
	IsEdited  bool      `json:"isEdited,omitempty"`
	Message   string    `json:"message"`
	Parent    int       `json:"parent,omitempty"`
	ThreadId  int       `json:"thread,omitempty"`
	Thread    string    `json:"-"`
	Path      string    `json:"-"`
	Childrens int       `json:"-"`
}

// ключ = types.Post.Id
type PostConnectionsItem struct {
	ThreadId         int    // types.Post.ThreadId
	MaterializedPath string // types.Post.MaterializedPath
	NumberOfChildren int    // types.Post.NumberOfChildren
}

//easyjson:json
type PostDetails struct {
	Author *User   `json:"author,omitempty"`
	Forum  *Forum  `json:"forum,omitempty"`
	Post   *Post   `json:"post"`
	Thread *Thread `json:"thread,omitempty"`
}
