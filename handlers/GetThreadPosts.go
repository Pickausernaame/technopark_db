package handlers

import (
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) GetThreadPosts(c *gin.Context) {
	slug_or_id := c.Param("slug_or_id")
	id, err := strconv.Atoi(slug_or_id)
	thread := &models.Thread{}
	if err != nil {
		thread, err = h.Agregator.GetThreadAgr(slug_or_id)
		if err != nil {
			c.JSON(404, err)
			return
		}
	} else {
		thread, err = h.Agregator.GetThreadById(id)
		if err != nil {
			c.JSON(404, err)
			return
		}
	}

	lim, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		lim = 100
	}

	since, err := strconv.Atoi(c.Query("since"))
	if err != nil {
		since = -1
	}
	desc, _ := strconv.ParseBool(c.Query("desc"))
	err = nil
	posts := &models.Posts{}
	switch c.Query("sort") {
	case "":
		fallthrough
	case "flat":
		if desc {
			posts, err = h.Agregator.GetPostsFlatDESC(thread.Id, lim, since)
		} else {
			posts, err = h.Agregator.GetPostsFlatASC(thread.Id, lim, since)
		}
	case "tree":
		if desc {
			posts, err = h.Agregator.GetPostsTreeDESC(thread.Id, lim, since)
		} else {
			posts, err = h.Agregator.GetPostsTreeASC(thread.Id, lim, since)
		}
	case "parent_tree":
		if desc {
			posts, err = h.Agregator.GetPostsParentTreeDESC(thread.Id, lim, since)
		} else {
			posts, err = h.Agregator.GetPostsParentTreeASC(thread.Id, lim, since)
		}
	}
	if err != nil {
		c.JSON(404, err)
		return
	}

	if len(*posts) == 0 {
		emptyArray := make([]int64, 0)
		c.JSON(200, emptyArray)
		return
	}
	c.JSON(200, posts)
}
