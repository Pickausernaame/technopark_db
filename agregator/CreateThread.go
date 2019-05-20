package agregator

import (
	"technopark_db/models"
)

func (agr *Agregator) CreateThreadAgr(thread models.Thread) (resThread models.Thread, err error) {
	sql := `INSERT INTO thread (author, created, forum, message, slug, title)
				VALUES ((SELECT nickname FROM users WHERE nickname = $1), $2,
						(SELECT slug FROM forum WHERE slug = $3), $4, $5, $6) 
				RETURNING author, created, forum, id, message, slug, title, votes;`
	err = agr.Connection.QueryRow(sql, thread.Author, thread.Created, thread.Forum, thread.Message, thread.Slug,
		thread.Title).Scan(&resThread.Author, &resThread.Created, &resThread.Forum, &resThread.Id, &resThread.Message, &resThread.Slug, &resThread.Title, &resThread.Votes)
	return
}
