package agregator

import (
	"fmt"
	"strconv"
	"strings"
	"technopark_db/models"
)

func (agr *Agregator) CreatePostByIdAgr(Posts []models.Post, id int) (outPosts []models.Post, err error) {
	sql := `
			INSERT INTO post (author, forum, thread_id, created, isEdited, message)
					VALUES ($1, $2, $3, $4, $5, $6)
					RETURNING id;`
	thread, err := agr.GetThreadById(id)
	if err != nil {
		return nil, err
	}
	for _, p := range Posts {

		fmt.Println(p.Parent)

		p.Forum = thread.Forum
		p.ThreadId = thread.Id
		p.Thread = thread.Slug
		err = agr.Connection.QueryRow(sql, p.Author, p.Forum, p.ThreadId, p.Created, p.IsEdited, p.Message).Scan(&p.Id)
		if err != nil {
			fmt.Println("Ошибка тут")
			return nil, err
		}
		outPosts = append(outPosts, p)
	}
	return
}

func (agr *Agregator) CreatePostBySlugAgr(Posts []models.Post, slug string) (outPosts []models.Post, err error) {
	sql := `
			INSERT INTO post (author, forum, thread_id, created, isEdited, message)
					VALUES ($1, $2, $3, $4, $5, $6)
					RETURNING id;`
	thread, err := agr.GetThreadAgr(slug)
	if err != nil {
		return nil, err
	}
	for _, p := range Posts {

		fmt.Println(p.Parent)
		p.Forum = thread.Forum
		p.ThreadId = thread.Id
		p.Thread = thread.Slug
		err = agr.Connection.QueryRow(sql, p.Author, p.Forum, p.ThreadId, p.Created, p.IsEdited, p.Message).Scan(&p.Id)
		if err != nil {
			fmt.Println("Ошибка тут")
			return nil, err
		}
		outPosts = append(outPosts, p)
	}
	return
}

type PostConnections map[int]models.PostConnectionsItem

func (agr *Agregator) GetPostConnections(parents []int) (postConnections PostConnections, err error) {
	parentsAsString := make([]string, len(parents))
	sql := `
			SELECT id, thread_id, path, children FROM post
					WHERE id IN (`
	for i, n := range parents {
		parentsAsString[i] = strconv.Itoa(n)
	}
	sql = sql + strings.Join(parentsAsString, ",") + ");"
	fmt.Println(sql)

	rows, err := agr.Connection.Query(sql)
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

func (agr *Agregator) PostsCreateInsert(posts []models.Post, roots int) (err error) {
	sql := ` 
	UPDATE thread SET roots = $1
		WHERE id = $2;`
	_, err = agr.Connection.Exec(sql, roots, posts[0].Id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//sql = `SELECT nextval('post_id_seq');`
	//var start int
	//err = agr.Connection.QueryRow(sql).Scan(&start)
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	for i := range posts {
		//oldId := posts[i].Id
		//posts[i].Id = start + i
		sql = `
			UPDATE post SET 
				path = $2, children = $3
			WHERE id = $1;`
		_, err = agr.Connection.Exec(sql, posts[i].Id, posts[i].Path, posts[i].Path)
		if err != nil {
			return err
		}
	}
	return
}
