package agregator

import "fmt"

func (agr *Agregator) ClearTableAgr() {
	sql := `TRUNCATE users, forum, thread, vote, post, usersforum RESTART IDENTITY CASCADE ;`
	_, err := agr.Connection.Exec(sql)
	fmt.Println(err)
}
