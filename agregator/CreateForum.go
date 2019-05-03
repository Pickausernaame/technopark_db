package agregator

import (
	"technopark_db/models"
)

func (agr *Agregator) CreateForumAgr(forum *models.Forum) (outForum models.Forum, err error) {
	sql := `
	INSERT INTO forum (slug, title, nickname)
		VALUES ( $1, $2, (SELECT nickname FROM users
								WHERE users.nickname = $3))
		RETURNING slug, title, nickname, posts, threads;`
	err = agr.Connection.QueryRow(sql, forum.Slug, forum.Title, forum.User).Scan(&outForum.Slug, &outForum.Title, &outForum.User, &outForum.Posts, &outForum.Threads)
	return
}

func (agr *Agregator) ErrorCreateForumAgr(slug string) (outForum models.Forum, err error) {
	sql := `
	SELECT slug, title, nickname, posts, threads FROM forum
		WHERE slug = $1;
	`
	err = agr.Connection.QueryRow(sql, slug).Scan(&outForum.Slug, &outForum.Title, &outForum.User, &outForum.Posts, &outForum.Threads)
	return
}
