package agregator

import "technopark_db/models"

func (agr *Agregator) ThreadUpdate(thread models.Thread) (err error) {
	sql := "UPDATE thread SET message = $2, title = $3 WHERE id = $1;"
	_, err = agr.Connection.Exec(sql, thread.Id, thread.Message, thread.Title)
	return
}
