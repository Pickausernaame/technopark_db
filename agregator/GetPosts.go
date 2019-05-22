package agregator

import (
	"fmt"
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/jackc/pgx"
	"strconv"
	"strings"
)

func NodeSetter(numb int) string {
	path := strconv.FormatInt(int64(numb), 10)
	return strings.Repeat("0", 6-len(path)) + path
}

func (agr *Agregator) GetPostsFlatASC(id int, lim int, since int) (outPosts []models.Post, err error) {
	var sql string
	var rows *pgx.Rows
	if since < 0 {
		sql = `
			SELECT id, author, created , is_edited, message, path, parent, children, forum FROM post
				WHERE thread_id  = $1
				ORDER BY (thread_id, id) ASC LIMIT $2;`
		rows, err = agr.Connection.Query(sql, id, lim)
	} else {

		sql = `
			SELECT id, author, created, is_edited, message, path, parent, children, forum FROM post
					WHERE thread_id  = $1 AND id > $2
					ORDER BY (thread_id, id) ASC LIMIT $3;`
		rows, err = agr.Connection.Query(sql, id, since, lim)

	}
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p models.Post

		err = rows.Scan(&p.Id, &p.Author, &p.Created, &p.IsEdited, &p.Message, &p.Path, &p.Parent, &p.Childrens, &p.Forum)
		p.ThreadId = id
		outPosts = append(outPosts, p)
	}
	return
}

func (agr *Agregator) GetPostsFlatDESC(id int, lim int, since int) (outPosts []models.Post, err error) {
	var sql string
	var rows *pgx.Rows
	if since < 0 {
		sql = `
				SELECT id, author, created, is_edited, message, path, parent, children, forum FROM post
					WHERE thread_id  = $1
					ORDER BY (thread_id, id) DESC LIMIT $2;`
		rows, err = agr.Connection.Query(sql, id, lim)
	} else {
		sql = `
				SELECT id, author, created, is_edited, message, path, parent, children, forum FROM post
					WHERE thread_id  = $1 AND id < $2
					ORDER BY (thread_id, id) DESC LIMIT $3;`
		rows, err = agr.Connection.Query(sql, id, since, lim)
	}
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p models.Post
		err = rows.Scan(&p.Id, &p.Author, &p.Created, &p.IsEdited, &p.Message, &p.Path, &p.Parent, &p.Childrens, &p.Forum)
		p.ThreadId = id
		outPosts = append(outPosts, p)
	}
	return
}

func (agr *Agregator) GetPostsTreeASC(id int, lim int, since int) (outPosts []models.Post, err error) {
	var sql string
	var rows *pgx.Rows
	if since < 0 {
		sql = `
			SELECT id, author,created, is_edited, message, path, parent, children, forum FROM post 
				WHERE thread_id  = $1 ORDER BY (thread_id, path) ASC LIMIT $2;`
		rows, err = agr.Connection.Query(sql, id, lim)
	} else {
		sql = `
			SELECT id, author,created , is_edited, message, path, parent, children, forum FROM post
					WHERE thread_id  = $1 AND path > (SELECT path FROM post WHERE id = $2)
					ORDER BY (thread_id, path) ASC LIMIT $3;`
		rows, err = agr.Connection.Query(sql, id, since, lim)
	}
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p models.Post
		p.ThreadId = id
		err = rows.Scan(&p.Id, &p.Author, &p.Created, &p.IsEdited, &p.Message, &p.Path, &p.Parent, &p.Childrens, &p.Forum)
		outPosts = append(outPosts, p)
	}
	return
}

func (agr *Agregator) GetPostsTreeDESC(id int, lim int, since int) (outPosts []models.Post, err error) {
	var sql string
	var rows *pgx.Rows
	if since < 0 {
		sql = `
				SELECT id, author, created , is_edited, message, path, parent, children, forum FROM post
						WHERE thread_id  = $1
						ORDER BY (thread_id, path) DESC LIMIT $2;`
		rows, err = agr.Connection.Query(sql, id, lim)
	} else {
		sql = `
				SELECT id, author, created, is_edited, message, path, parent, children, forum FROM post
						WHERE thread_id  = $1 AND path < (SELECT path FROM post WHERE id = $2)
						ORDER BY (thread_id, path) DESC LIMIT $3;`
		rows, err = agr.Connection.Query(sql, id, since, lim)
	}
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p models.Post
		p.ThreadId = id
		err = rows.Scan(&p.Id, &p.Author, &p.Created, &p.IsEdited, &p.Message, &p.Path, &p.Parent, &p.Childrens, &p.Forum)
		outPosts = append(outPosts, p)
	}
	return
}

func (agr *Agregator) GetPostsParentTreeASC(id int, lim int, since int) (outPosts []models.Post, err error) {
	var sql string
	var rows *pgx.Rows
	if since < 0 {
		sql = `
			SELECT id, author, created, is_edited, message, path, parent, children, forum FROM post
					WHERE thread_id  = $1 AND path > '000000' AND path < $2
					ORDER BY (thread_id, path) ASC;`
		pathLim := NodeSetter(lim + 1)
		rows, err = agr.Connection.Query(sql, id, pathLim)
	} else {
		var path string
		sql = `SELECT path FROM post WHERE id = $1;`
		err = agr.Connection.QueryRow(sql, since).Scan(&path)
		if err != nil {
			fmt.Println()
			return
		}
		sql = `
			SELECT id, author, created , is_edited, message, path, parent, children, forum FROM post
					WHERE thread_id  = $1 AND path > $2 AND path < $3
					ORDER BY (thread_id, path) ASC;`
		pathFrom, _ := strconv.Atoi(path[0:6])
		pathLim := NodeSetter(lim + 1 + pathFrom)
		rows, err = agr.Connection.Query(sql, id, path, pathLim)
	}
	if err != nil {

		return
	}

	defer rows.Close()
	for rows.Next() {
		var p models.Post
		p.ThreadId = id
		err = rows.Scan(&p.Id, &p.Author, &p.Created, &p.IsEdited, &p.Message, &p.Path, &p.Parent, &p.Childrens, &p.Forum)
		outPosts = append(outPosts, p)
	}
	return
}

func (agr *Agregator) GetPostsParentTreeDESC(id int, lim int, since int) (outPosts []models.Post, err error) {
	sql := `WITH bounds AS (SELECT
  					SUBSTRING($2 FROM 1 FOR 6	) AS from_prefix,
					SUBSTRING($2 FROM 8			) AS from_postfix,
  					SUBSTRING($3 FROM 1 FOR 6	) AS to_prefix,
					SUBSTRING($3 FROM 8			) AS to_postfix
					)( SELECT id, author, created , is_edited, message, path, parent, children, forum FROM post, bounds
							WHERE	thread_id = $1	AND	SUBSTRING(path FROM 1 FOR 6) = bounds.from_prefix AND SUBSTRING(path FROM 8) < bounds.from_postfix
							ORDER BY thread_id, SUBSTRING(path FROM 1 FOR 6) DESC, SUBSTRING(path FROM 8) ASC)
						UNION ALL (SELECT id, author, created, is_edited, message, path, parent, children, forum FROM post, bounds
							WHERE	thread_id = $1 AND SUBSTRING(path FROM 1 FOR 6) > bounds.from_prefix AND SUBSTRING(path FROM 1 FOR 6) < bounds.to_prefix
							ORDER BY thread_id, SUBSTRING(path FROM 1 FOR 6) DESC, SUBSTRING(path FROM 8) ASC)
						UNION ALL (SELECT id, author, created, is_edited, message, path, parent, children, forum FROM post, bounds
							WHERE	thread_id = $1 AND SUBSTRING(path FROM 1 FOR 6) = bounds.to_prefix AND SUBSTRING(path from 8) > bounds.to_postfix
							ORDER BY SUBSTRING(path FROM 1 FOR 6) DESC, SUBSTRING(path FROM 8) ASC);`

	var rows *pgx.Rows
	if since < 0 {

		ssql := `SELECT roots FROM thread WHERE id = $1;`
		var pathDown int
		err = agr.Connection.QueryRow(ssql, id).Scan(&pathDown)
		if err != nil {
			fmt.Println(err)
			return
		}
		mPathDown := NodeSetter(pathDown + 1)
		pathUp := pathDown - lim
		if pathUp < 0 {
			pathUp = 0
		}
		mPathUp := NodeSetter(pathUp)
		rows, err = agr.Connection.Query(sql, id, mPathUp, mPathDown)
	} else {
		var mPathDown string
		ssql := `SELECT path FROM post WHERE id = $1;`
		err = agr.Connection.QueryRow(ssql, since).Scan(&mPathDown)
		if err != nil {
			return
		}
		pathDown, _ := strconv.Atoi(mPathDown[0:6])
		pathUp := pathDown - lim - 1
		if pathUp < 0 {
			pathUp = 0
		}
		mPathUp := NodeSetter(pathUp)
		rows, err = agr.Connection.Query(sql, id, mPathUp, mPathDown)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p models.Post
		p.ThreadId = id
		err = rows.Scan(&p.Id, &p.Author, &p.Created, &p.IsEdited, &p.Message, &p.Path, &p.Parent, &p.Childrens, &p.Forum)
		outPosts = append(outPosts, p)
		fmt.Println(p.Path)
	}
	return
}
