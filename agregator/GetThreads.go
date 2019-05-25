package agregator

import (
	"fmt"
	"github.com/Pickausernaame/technopark_db/models"
	"time"
)

func (agr *Agregator) GetThreadsAgr(slug string, limit int, since time.Time, desk bool) (threads *[]models.Thread, err error, exist bool) {
	sql := `SELECT EXISTS (SELECT true FROM forum WHERE slug = $1);`
	//sql := `SELECT * FROM forum WHERE slug = $1;`
	err = agr.Connection.QueryRow(sql, slug).Scan(&exist)
	if !exist {
		return nil, err, false
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
	if err != nil {
		return nil, err, false
	}
	threads = &[]models.Thread{}
	for rows.Next() {
		var thread models.Thread
		err = rows.Scan(&thread.Author, &thread.Created, &thread.Forum, &thread.Id, &thread.Message, &thread.Slug, &thread.Title, &thread.Votes)
		*threads = append(*threads, thread)
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}
