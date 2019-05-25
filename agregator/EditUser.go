package agregator

import (
	"github.com/Pickausernaame/technopark_db/models"
)

func (agr *Agregator) EditUserAgr(nickname string, editUser *models.User) (curUser *models.User, err error) {
	sql := `UPDATE users SET 
				nickname = CASE WHEN $2 = '' THEN nickname ELSE $2 END,
				fullname = CASE WHEN $3 = '' THEN fullname ELSE $3 END,
				about = CASE WHEN $4 = '' THEN about ELSE $4 END,
				email = CASE WHEN $5 = '' THEN email ELSE $5 END
				
			WHERE nickname = $1
			RETURNING nickname, fullname, about, email;
`
	curUser = &models.User{}
	err = agr.Connection.QueryRow(sql, nickname, editUser.Nickname, editUser.Fullname, editUser.About,
		editUser.Email).Scan(&curUser.Nickname, &curUser.Fullname, &curUser.About, &curUser.Email)

	return
}
