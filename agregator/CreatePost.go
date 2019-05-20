package agregator

import (
	"fmt"
	"strconv"
	"strings"
	"technopark_db/models"
)

func (agr *Agregator) CreatePostByIdAgr(Posts []models.Post, id int, created bool) (outPosts []models.Post, err error) {
	tx, err := agr.Connection.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		return
	}()
	thread, err := agr.GetThreadById(id)
	if err != nil {
		fmt.Println("Ошибка чиения треда по ID")
		return nil, err
	}

	sql := ``
	if created {
		sql = `INSERT INTO post (author, forum, thread_id, is_edited, message, parent, created)
					VALUES ($1, $2, $3, $4, $5, $6, $7)
					RETURNING id;`
		for _, p := range Posts {
			p.Forum = thread.Forum
			p.ThreadId = thread.Id
			p.Thread = thread.Slug
			err = tx.QueryRow(sql, p.Author, p.Forum, p.ThreadId, p.IsEdited, p.Message, p.Parent, p.Created).Scan(&p.Id)
			if err != nil {
				fmt.Println("Ошибка создания поста по slug")
				fmt.Println(err)
				return nil, err
			}
			outPosts = append(outPosts, p)
		}
	} else {
		sql = `INSERT INTO post (author, forum, thread_id, is_edited, message, parent)
					VALUES ($1, $2, $3, $4, $5, $6)
					RETURNING id, created;`
		for _, p := range Posts {
			p.Forum = thread.Forum
			p.ThreadId = thread.Id
			p.Thread = thread.Slug
			err = tx.QueryRow(sql, p.Author, p.Forum, p.ThreadId, p.IsEdited, p.Message, p.Parent).Scan(&p.Id, &p.Created)
			if err != nil {
				fmt.Println("Ошибка создания поста по slug")
				fmt.Println(err)
				return nil, err
			}
			outPosts = append(outPosts, p)
		}
	}
	return
}

func (agr *Agregator) CreatePostBySlugAgr(Posts []models.Post, slug string, created bool) (outPosts []models.Post, err error) {
	tx, err := agr.Connection.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		return
	}()
	thread, err := agr.GetThreadAgr(slug)
	if err != nil {
		fmt.Println("Ошибка чиения треда по slug")
		fmt.Println(err)
		return nil, err
	}
	sql := ``
	if created {
		sql = `INSERT INTO post (author, forum, thread_id, is_edited, message, parent, created)
					VALUES ($1, $2, $3, $4, $5, $6, $7)
					RETURNING id;`
		for _, p := range Posts {
			p.Forum = thread.Forum
			p.ThreadId = thread.Id
			p.Thread = thread.Slug
			err = tx.QueryRow(sql, p.Author, p.Forum, p.ThreadId, p.IsEdited, p.Message, p.Parent, p.Created).Scan(&p.Id)
			if err != nil {
				fmt.Println("Ошибка создания поста по slug")
				fmt.Println(err)
				return nil, err
			}
			outPosts = append(outPosts, p)
		}
	} else {
		sql = `INSERT INTO post (author, forum, thread_id, is_edited, message, parent)
					VALUES ($1, $2, $3, $4, $5, $6)
					RETURNING id, created;`
		for _, p := range Posts {
			p.Forum = thread.Forum
			p.ThreadId = thread.Id
			p.Thread = thread.Slug
			err = tx.QueryRow(sql, p.Author, p.Forum, p.ThreadId, p.IsEdited, p.Message, p.Parent).Scan(&p.Id, &p.Created)
			if err != nil {
				fmt.Println("Ошибка создания поста по slug")
				fmt.Println(err)
				return nil, err
			}
			outPosts = append(outPosts, p)
		}
	}
	return
}

func (agr *Agregator) GetRoots(id int) (roots int, err error) {
	sql := `SELECT roots FROM thread WHERE id = $1;`
	err = agr.Connection.QueryRow(sql, id).Scan(&roots)
	return
}

type PostConnections map[int]models.PostConnectionsItem

func (agr *Agregator) GetPostConnections(parents []int) (postConnections PostConnections, err error) {
	tx, err := agr.Connection.Begin()
	parentsAsString := make([]string, len(parents))
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		return
	}()

	sql := `
			SELECT id, thread_id, path, children FROM post
					WHERE id IN (`
	for i, n := range parents {
		parentsAsString[i] = strconv.Itoa(n)
	}
	sql = sql + strings.Join(parentsAsString, ",") + ");"
	rows, err := tx.Query(sql)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	postConnections = make(PostConnections)
	for rows.Next() {
		var id int
		var item models.PostConnectionsItem
		err = rows.Scan(&id, &item.ThreadId, &item.MaterializedPath, &item.NumberOfChildren)
		if err != nil {
			return nil, err
		}
		postConnections[id] = item
	}
	return postConnections, nil
}

func (agr *Agregator) PostChildrenUpdate(pc PostConnections) (err error) {
	tx, err := agr.Connection.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		return
	}()
	sql := "UPDATE post set children = $2 WHERE id = $1;"
	for k, v := range pc {
		_, err = tx.Exec(sql, k, v.NumberOfChildren)
	}
	return
}

func (agr *Agregator) PostsCreateInsert(posts []models.Post, roots int) (err error) {

	tx, err := agr.Connection.Begin()

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		return
	}()

	if roots != 0 {
		sql := ` 
	UPDATE thread SET roots = $1
		WHERE id = $2;`
		_, err = tx.Exec(sql, roots, posts[0].ThreadId)
		if err != nil {
			fmt.Println("Ошибка обновления корней")
			fmt.Println(err)
			return err
		}
	}

	for _, p := range posts {

		sql := `
			UPDATE post SET 
				path = $2, children = $3
			WHERE id = $1;`
		_, err = tx.Exec(sql, p.Id, p.Path, p.Childrens)
		if err != nil {
			fmt.Println("Ошибка обновления детей")
			return err
		}

	}
	return
}

func (agr *Agregator) InsertUsersInForum(posts []models.Post) (err error) {
	tx, err := agr.Connection.Begin()

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		return
	}()

	for _, p := range posts {
		ssql := `INSERT INTO usersforum (nickname, forum) VALUES ($1, $2)
					ON CONFLICT DO NOTHING;`
		_, err = tx.Exec(ssql, p.Author, p.Forum)
		if err != nil {
			fmt.Println("Ошибка обновления usersforum", err)
			return err
		}
	}
	return

}
