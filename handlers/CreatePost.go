package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Pickausernaame/technopark_db/agregator"
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func NodeSetter(numb int) string {
	path := strconv.FormatInt(int64(numb), 10)
	return strings.Repeat("0", 6-len(path)) + path
}

func (h *Handler) CreatePost(c *gin.Context) {
	var posts *models.Posts
	err := json.NewDecoder(c.Request.Body).Decode(&posts)
	buf, _ := c.GetRawData()
	created := strings.Contains(string(buf), "created")
	slug_or_id := c.Param("slug_or_id")
	id, err := strconv.Atoi(slug_or_id)

	if err != nil {
		exist, err := h.Agregator.IsThreadExistBySlug(slug_or_id)
		if !exist || err != nil {
			c.JSON(404, err)
			return
		}
		if len(*posts) == 0 {
			emptyArray := make([]int64, 0)
			c.JSON(201, emptyArray)
			return
		}
		thread, err := h.Agregator.GetThreadAgr(slug_or_id)
		if err != nil {
			c.JSON(404, err)
			return
		}
		for _, p := range *posts {
			p.ThreadId = thread.Id
			p.Forum = thread.Forum
		}
		posts, err = h.Agregator.CreatePostBySlugAgr(posts, slug_or_id, created)
		if err != nil {
			c.JSON(404, err)
			return
		}
	} else {
		exist, err := h.Agregator.IsThreadExistById(id)
		if !exist || err != nil {
			c.JSON(404, err)
			return
		}
		if len(*posts) == 0 {
			emptyArray := make([]int64, 0)
			c.JSON(201, emptyArray)
			return
		}
		posts, err = h.Agregator.CreatePostByIdAgr(posts, id, created)
		if err != nil {
			c.JSON(404, err)
			return
		}
	}

	var postsConnections agregator.PostConnections
	parentIds := make([]int, len(*posts))
	for i := range *posts {
		if (*posts)[i].Parent != 0 {
			parentIds = append(parentIds, (*posts)[i].Parent)
		}

	}
	postsConnections, err = h.Agregator.GetPostConnections(parentIds)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, err)
		return
	}

	NumberOfRoots, err := h.Agregator.GetRoots((*posts)[0].ThreadId)
	for i := range *posts {
		if (*posts)[i].Parent == 0 {
			NumberOfRoots++
			(*posts)[i].Path = NodeSetter(NumberOfRoots)
		} else {
			parentId := (*posts)[i].Parent

			parentPostConnection := postsConnections[parentId]

			if parentPostConnection.ThreadId != (*posts)[i].ThreadId {
				c.JSON(409, models.Thread{})
				return
			}
			parentPostConnection.NumberOfChildren++
			(*posts)[i].Path = parentPostConnection.MaterializedPath + "." + NodeSetter(parentPostConnection.NumberOfChildren)
			postsConnections[parentId] = parentPostConnection
		}
	}

	err = h.Agregator.PostChildrenUpdate(&postsConnections)
	if err != nil {
		c.JSON(404, err)
		return
	}
	err = h.Agregator.PostsCreateInsert(posts, NumberOfRoots)

	if err != nil {
		fmt.Println(err)

		c.JSON(404, err)
		return
	}
	err = h.Agregator.InsertUsersInForum(posts)
	c.JSON(201, posts)
}
