package agregator

import "github.com/Pickausernaame/technopark_db/models"

func (agr *Agregator) Status() (status *models.Service, err error) {
	// Todo: навесить триггер
	sql := `
	SELECT
		(SELECT count(*) FROM users)	AS count_user,
		(SELECT count(*) FROM forum)	AS count_forum,
		(SELECT count(*) FROM thread)	AS count_thread,
		(SELECT count(*) FROM post)		AS count_post;`
	status = &models.Service{}
	err = agr.Connection.QueryRow(sql).Scan(&status.Users, &status.Forums, &status.Threads, &status.Posts)
	return
}
