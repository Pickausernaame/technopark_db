package agregator

import (
	"fmt"
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/jackc/pgx"
	"strconv"
	"strings"
)

func generateBody(len int, argc int) (body string) {
	format := "("
	for i := 0; i < argc-1; i++ {
		format += "$%d,"
	}
	format += "$%d)"
	for i := 0; i < len; i++ {
		var argv []interface{}
		for j := 0; j < argc; j++ {
			num := i*argc + j + 1
			argv = append(argv, num)
		}
		part := fmt.Sprintf(format, argv...)
		if i < len-1 {
			part += ","
		}
		body += part

	}
	return body
}

func (agr *Agregator) InsertWithCreated(posts *models.Posts, thread *models.Thread, tx *pgx.Tx) (outPosts *models.Posts, err error) {
	const argc = 7
	sql := "INSERT INTO post (author, forum, thread_id, is_edited, message, parent, created) VALUES " + generateBody(len(*posts), argc) + "RETURNING id;"
	var argv []interface{}
	outPosts = &models.Posts{}
	for _, p := range *posts {
		p.Forum = thread.Forum
		p.ThreadId = thread.Id
		p.Thread = thread.Slug
		argv = append(argv, p.Author, p.Forum, p.ThreadId, p.IsEdited, p.Message, p.Parent, p.Created)
		*outPosts = append(*outPosts, p)
	}
	rows, err := tx.Query(sql, argv...)
	if err != nil {
		fmt.Println("Ошибка создания поста по slug")
		return nil, err
	}
	ind := -1
	for rows.Next() {
		ind++
		err = rows.Scan(&(*outPosts)[ind].Id)
		if err != nil {
			return nil, err
		}
	}
	return
}

func (agr *Agregator) InsertWithoutCreated(posts *models.Posts, thread *models.Thread, tx *pgx.Tx) (outPosts *models.Posts, err error) {
	const argc = 6
	fmt.Println(generateBody(len(*posts), argc))
	sql := "INSERT INTO post (author, forum, thread_id, is_edited, message, parent) VALUES " + generateBody(len(*posts), argc) + "RETURNING id, created;"
	var argv []interface{}
	outPosts = &models.Posts{}
	for _, p := range *posts {
		p.Forum = thread.Forum
		p.ThreadId = thread.Id
		p.Thread = thread.Slug
		argv = append(argv, p.Author, p.Forum, p.ThreadId, p.IsEdited, p.Message, p.Parent)

		*outPosts = append(*outPosts, p)
	}

	rows, err := tx.Query(sql, argv...)
	if err != nil {
		fmt.Println("Ошибка создания поста по slug")
		fmt.Println(err)
		return nil, err
	}
	ind := -1

	for rows.Next() {
		ind++
		err := rows.Scan(&(*outPosts)[ind].Id, &(*outPosts)[ind].Created)
		if err != nil {
			return nil, err
		}
	}
	return
}

func (agr *Agregator) CreatePostByIdAgr(posts *models.Posts, id int, created bool) (outPosts *models.Posts, err error) {
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

	if created {
		outPosts, err = agr.InsertWithCreated(posts, thread, tx)
		if err != nil {
			return nil, err
		}
	} else {
		outPosts, err = agr.InsertWithoutCreated(posts, thread, tx)
		if err != nil {
			return nil, err
		}
	}
	return
}

func (agr *Agregator) CreatePostBySlugAgr(posts *models.Posts, slug string, created bool) (outPosts *models.Posts, err error) {
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
		return nil, err
	}
	if created {
		outPosts, err = agr.InsertWithCreated(posts, thread, tx)
		if err != nil {
			return nil, err
		}
	} else {
		outPosts, err = agr.InsertWithoutCreated(posts, thread, tx)
		if err != nil {
			return nil, err
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

func (agr *Agregator) PostChildrenUpdate(pc *PostConnections) (err error) {
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
	for k, v := range *pc {
		_, err = tx.Exec(sql, k, v.NumberOfChildren)
	}
	return
}

func (agr *Agregator) PostsCreateInsert(posts *models.Posts, roots int) (err error) {

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
		_, err = tx.Exec(sql, roots, (*posts)[0].ThreadId)
		if err != nil {
			fmt.Println("Ошибка обновления корней")
			fmt.Println(err)
			return err
		}
	}

	for _, p := range *posts {
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

func (agr *Agregator) InsertUsersInForum(posts *models.Posts) (err error) {
	tx, err := agr.Connection.Begin()
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
		return
	}()

	for _, p := range *posts {
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
