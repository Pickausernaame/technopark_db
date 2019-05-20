package agregator

import (
	"fmt"
	"technopark_db/models"
	"time"
)

func (agr *Agregator) GetThreadsAgr(slug string, limit int, since time.Time, desk bool) (Threads []models.Thread, err error, exist bool) {
	sql := `SELECT EXISTS (SELECT true FROM forum WHERE slug = $1);`
	//sql := `SELECT * FROM forum WHERE slug = $1;`
	err = agr.Connection.QueryRow(sql, slug).Scan(&exist)
	if !exist {
		return
	}
	if desk {
		sql = `
				SELECT author, created, forum, id, message, slug, title, votes FROM thread
					WHERE forum = $1 AND created <= $3
					ORDER BY created DESC LIMIT $2;`
	} else {
		sql = `
				SELECT author, created, forum, id, message, slug, title, votes FROM thread
					WHERE forum = $1 AND created >= $3
					ORDER BY created ASC LIMIT $2;`
	}
	rows, err := agr.Connection.Query(sql, slug, limit, since)

	for rows.Next() {
		var thread models.Thread
		err = rows.Scan(&thread.Author, &thread.Created, &thread.Forum, &thread.Id, &thread.Message, &thread.Slug, &thread.Title, &thread.Votes)
		Threads = append(Threads, thread)
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}
