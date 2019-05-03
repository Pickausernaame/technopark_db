package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"technopark_db/agregator"
	"technopark_db/models"
)

func NodeSetter(numb int) string {
	path := strconv.FormatInt(int64(numb), 10)
	return strings.Repeat("0", 6-len(path)) + path
}

func (h *Handler) CreatePost(c *gin.Context) {
	var posts []models.Post
	_ = json.NewDecoder(c.Request.Body).Decode(&posts)
	if len(posts) == 0 {
		emptyArray := make([]int64, 0)
		c.JSON(201, emptyArray)
		return
	}
	slug_or_id := c.Param("slug_or_id")
	id, err := strconv.Atoi(slug_or_id)
	var outPosts []models.Post
	if err != nil {
		thread, _ := h.Agregator.GetThreadAgr(slug_or_id)
		for _, p := range posts {

			p.ThreadId = thread.Id
			p.Forum = thread.Forum
		}
		posts, err = h.Agregator.CreatePostBySlugAgr(posts, slug_or_id)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, err)
			return
		}
		//c.JSON(201, outPosts)
	} else {
		posts, err = h.Agregator.CreatePostByIdAgr(posts, id)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, err)
			return
		}
		//c.JSON(201, outPosts)
	}
	var postsConnections agregator.PostConnections
	parentIds := make([]int, len(posts))
	for i := range outPosts {
		if outPosts[i].Parent != 0 {
			parentIds = append(parentIds, posts[i].Parent)
		}
	}
	postsConnections, err = h.Agregator.GetPostConnections(parentIds)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, err)
		return
	}
	var NumberOfRoots int
	for i := range posts {
		if posts[i].Parent == 0 {
			NumberOfRoots++
			posts[i].Path = NodeSetter(NumberOfRoots)
		} else { // если родитель - post
			parentId := posts[i].Parent
			parentPostConnections := postsConnections[parentId]
			fmt.Println(parentPostConnections.ThreadId, posts[i].ThreadId)
			if parentPostConnections.ThreadId != posts[i].ThreadId {
				fmt.Println(err)
				c.JSON(409, "error")
				return
			}
			parentPostConnections.NumberOfChildren++
			posts[i].Path = parentPostConnections.MaterializedPath + "." + NodeSetter(parentPostConnections.NumberOfChildren)
			postsConnections[parentId] = parentPostConnections
		}
	}
	for i := range postsConnections {
		for j := range posts {
			if i == posts[j].Id {
				posts[j].Childrens = postsConnections[i].NumberOfChildren
			}
		}
	}

	err = h.Agregator.PostsCreateInsert(posts, NumberOfRoots)

	if err != nil {
		fmt.Println(err)
		c.JSON(404, err)
		return
	}
	c.JSON(201, posts)
}
