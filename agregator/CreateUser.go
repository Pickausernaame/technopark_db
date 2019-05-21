package agregator

import (
	"fmt"
	"github.com/Pickausernaame/technopark_db/models"
)

func (agr *Agregator) CreateUserAgr(newUser models.User) (err error) {
	sql := `INSERT INTO users (nickname, fullname, about, email)
				VALUES($1, $2, $3, $4);`
	tx, err := agr.Connection.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		return
	}()

	_, err = tx.Exec(sql, newUser.Nickname, newUser.Fullname, newUser.About, newUser.Email)

	//err = agr.Connection.QueryRow(sql, newUser.Nickname, newUser.Fullname, newUser.About, newUser.Email).Scan(&resUser.Nickname, &resUser.Fullname, &resUser.About, &resUser.Email)
	if err != nil {
		fmt.Println("Error in agregation CreateUser")
		fmt.Println(err)
	}
	return
}

func (agr *Agregator) ErrorCreateUserArg(nicname string, email string) (resUsers []models.User, err error) {
	sql := `SELECT nickname, fullname, about, email FROM users
				WHERE nickname = $1 OR email = $2;`
	rows, err := agr.Connection.Query(sql, nicname, email)
	defer rows.Close()
	for rows.Next() {
		var currentUser models.User
		err = rows.Scan(&currentUser.Nickname, &currentUser.Fullname, &currentUser.About, &currentUser.Email)
		resUsers = append(resUsers, currentUser)
	}

	return
}
