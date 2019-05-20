package agregator

import (
	"fmt"
	"github.com/Pickausernaame/technopark_db/models"
)

func (agr *Agregator) GetForumAgr(slug string) (outForum models.Forum, err error) {
	sql := `SELECT slug, title, nickname, posts, threads FROM forum
				WHERE forum.slug = $1;`

	err = agr.Connection.QueryRow(sql, slug).Scan(&outForum.Slug, &outForum.Title, &outForum.User, &outForum.Posts, &outForum.Threads)
	fmt.Println(err)
	return
}

func (agr *Agregator) IsForumExist(slug string) (exist bool, err error) {
	sql := `SELECT EXISTS(SELECT true FROM forum WHERE slug = $1);`
	err = agr.Connection.QueryRow(sql, slug).Scan(&exist)
	return
}
