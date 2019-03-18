package agregator

import "fmt"

func (agr *Agregator) ClearTableAgr() {
	sql := `
			DROP TABLE IF EXISTS users CASCADE;
			DROP TABLE IF EXISTS forum CASCADE;
			DROP TABLE IF EXISTS thread CASCADE;
			DROP TABLE IF EXISTS post CASCADE;
			DROP TABLE IF EXISTS vote CASCADE;	
`

	_, err := agr.Connection.Exec(sql)
	fmt.Println(err)
}
