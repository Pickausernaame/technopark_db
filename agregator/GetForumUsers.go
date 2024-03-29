package agregator

import (
	"fmt"
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/jackc/pgx"
)

func (agr *Agregator) GetUsersASC(slug string, lim int, since string) (users *[]models.User, err error) {
	var rows *pgx.Rows

	if since != "" {
		//sql := `SELECT users.nickname, users.about, users.email, users.fullname FROM users
		//		INNER JOIN usersforum ON usersforum.nickname = users.nickname
		//		WHERE usersforum.forum = $1 AND usersforum.nickname > $2
		//			ORDER BY usersforum.nickname ASC LIMIT $3;`

		sql := `SELECT usersforum.nickname, users.about, users.email, users.fullname FROM usersforum, users
				WHERE usersforum.forum = $1 AND usersforum.nickname > $2 AND usersforum.nickname = users.nickname
					ORDER BY usersforum.nickname ASC LIMIT $3;`

		rows, err = agr.Connection.Query(sql, slug, since, lim)
	} else {
		//sql := `SELECT users.nickname, users.about, users.email, users.fullname FROM users
		//		INNER JOIN usersforum ON usersforum.nickname = users.nickname
		//		WHERE usersforum.forum = $1
		//			ORDER BY usersforum.nickname ASC LIMIT $2;`
		sql := `SELECT usersforum.nickname, users.about, users.email, users.fullname FROM usersforum, users
				WHERE usersforum.forum = $1 AND usersforum.nickname = users.nickname
					ORDER BY usersforum.nickname ASC LIMIT $2;`
		rows, err = agr.Connection.Query(sql, slug, lim)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	users = &[]models.User{}
	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.Nickname, &u.About, &u.Email, &u.Fullname)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		*users = append(*users, u)
	}
	return users, nil
}

func (agr *Agregator) GetUsersDESC(slug string, lim int, since string) (users *[]models.User, err error) {
	var rows *pgx.Rows
	if since != "" {
		sql := `SELECT usersforum.nickname, users.about, users.email, users.fullname FROM usersforum, users
				WHERE usersforum.forum = $1 AND usersforum.nickname = users.nickname AND usersforum.nickname < $2
					ORDER BY usersforum.nickname DESC LIMIT $3;`
		rows, err = agr.Connection.Query(sql, slug, since, lim)
	} else {
		sql := `SELECT usersforum.nickname, users.about, users.email, users.fullname FROM usersforum, users
				WHERE usersforum.forum = $1 AND usersforum.nickname = users.nickname
					ORDER BY usersforum.nickname DESC LIMIT $2;`
		rows, err = agr.Connection.Query(sql, slug, lim)
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var u models.User
	users = &[]models.User{}
	for rows.Next() {
		err = rows.Scan(&u.Nickname, &u.About, &u.Email, &u.Fullname)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		*users = append(*users, u)
	}
	return users, nil
}
