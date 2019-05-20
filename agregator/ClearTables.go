package agregator

func (agr *Agregator) ClearTableAgr() {
	sql := `TRUNCATE TABLE ONLY users, forum, thread, vote, post, usersforum RESTART IDENTITY;`
	agr.Connection.Exec(sql)
}
