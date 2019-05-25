package agregator

import (
	"github.com/Pickausernaame/technopark_db/models"
)

func (agr *Agregator) GetThreadAgr(slug string) (curThread *models.Thread, err error) {
	tx, err := agr.Connection.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		return
	}()
	sql := `
	SELECT author, created, forum, id, message, slug, title, votes FROM thread
		WHERE slug = $1;`
	curThread = &models.Thread{}
	err = tx.QueryRow(sql, slug).Scan(&curThread.Author, &curThread.Created, &curThread.Forum, &curThread.Id, &curThread.Message, &curThread.Slug, &curThread.Title, &curThread.Votes)
	return
}

func (agr *Agregator) GetThreadById(id int) (curThread *models.Thread, err error) {
	tx, err := agr.Connection.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		return
	}()
	sql := `
	SELECT author, created, forum, id, message, slug, title, votes FROM thread
		WHERE id = $1;`
	curThread = &models.Thread{}
	err = tx.QueryRow(sql, id).Scan(&curThread.Author, &curThread.Created, &curThread.Forum, &curThread.Id, &curThread.Message, &curThread.Slug, &curThread.Title, &curThread.Votes)
	return
}

func (agr *Agregator) IsThreadExistById(id int) (exist bool, err error) {
	sql := `SELECT EXISTS(SELECT true FROM thread WHERE id = $1);`
	err = agr.Connection.QueryRow(sql, id).Scan(&exist)
	return
}

func (agr *Agregator) IsThreadExistBySlug(slug string) (exist bool, err error) {
	sql := `SELECT EXISTS(SELECT true FROM thread WHERE slug = $1);`
	err = agr.Connection.QueryRow(sql, slug).Scan(&exist)
	return
}
