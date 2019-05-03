package agregator

import (
	"technopark_db/models"
)

func (agr *Agregator) GetThreadAgr(slug string) (curThread models.Thread, err error) {
	sql := `
	SELECT author, created, forum, id, message, slug, title, votes FROM thread
		WHERE slug = $1;`
	err = agr.Connection.QueryRow(sql, slug).Scan(&curThread.Author, &curThread.Created, &curThread.Forum, &curThread.Id, &curThread.Message, &curThread.Slug, &curThread.Title, &curThread.Votes)
	return
}

func (agr *Agregator) GetThreadById(id int) (curThread models.Thread, err error) {
	sql := `
	SELECT author, created, forum, id, message, slug, title, votes FROM thread
		WHERE id = $1;`
	err = agr.Connection.QueryRow(sql, id).Scan(&curThread.Author, &curThread.Created, &curThread.Forum, &curThread.Id, &curThread.Message, &curThread.Slug, &curThread.Title, &curThread.Votes)
	return
}
