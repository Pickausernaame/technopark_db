package agregator

import "technopark_db/models"

func (agr *Agregator) GetForumAgr(slug string) (outForum models.Forum, err error) {
	sql := `
	SELECT slug, title, nickname, posts, threads FROM forum
		WHERE slug = $1;
	`
	err = agr.Connection.QueryRow(sql, slug).Scan(&outForum.Slug, &outForum.Title, &outForum.User, &outForum.Posts, &outForum.Threads)
	return
}
