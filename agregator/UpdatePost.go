package agregator

func (agr *Agregator) PostUpdate(id int, message string) (err error) {
	sql := `UPDATE post SET message = $2, is_edited = true WHERE id = $1;`
	_, err = agr.Connection.Exec(sql, id, message)
	return
}
