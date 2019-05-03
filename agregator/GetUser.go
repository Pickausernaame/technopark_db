package agregator

import (
	"technopark_db/models"
)

func (agr *Agregator) GetUserAgr(nickname string) (curUser models.User, err error) {
	sql := `SELECT nickname, fullname, about, email  FROM users
				WHERE nickname = $1;`
	err = agr.Connection.QueryRow(sql, nickname).Scan(&curUser.Nickname, &curUser.Fullname, &curUser.About, &curUser.Email)
	//fmt.Println(err)
	return
}
