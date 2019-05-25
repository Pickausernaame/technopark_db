package agregator

func (agr *Agregator) Clastering() error {
	sql := `
		CLUSTER users USING nickname_ind;
		CLUSTER thread USING thread_created;
		CLUSTER forum USING forum_unique;
		CLUSTER post USING flat_sort;
`
	_, err := agr.Connection.Exec(sql)
	return err
}
