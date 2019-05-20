package agregator

import (
	"fmt"
	"technopark_db/models"
)

func (agr *Agregator) GetPost(id int) (post models.Post, err error) {
	sql := `SELECT author, forum, thread_id, created, is_edited, message, parent FROM post
				WHERE id = $1;`
	err = agr.Connection.QueryRow(sql, id).Scan(&post.Author, &post.Forum, &post.ThreadId, &post.Created, &post.IsEdited, &post.Message, &post.Parent)
	post.Id = id
	fmt.Println(post.Created.Local())
	return
}
