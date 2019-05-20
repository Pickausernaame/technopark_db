package handlers

import (
	"fmt"
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func (h *Handler) GetPostDetails(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, err)
		return
	}
	related := c.Query("related")
	needUser := strings.Contains(related, "user")
	needForum := strings.Contains(related, "forum")
	needThread := strings.Contains(related, "thread")

	var response models.PostDetails

	p, err := h.Agregator.GetPost(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, err)
		return
	}

	response.Post = &p

	if needUser {
		u, err := h.Agregator.GetUserAgr(p.Author)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, err)
			return
		}
		response.Author = &u
	}
	if needForum {
		f, err := h.Agregator.GetForumAgr(p.Forum)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, err)
			return
		}
		response.Forum = &f
	}
	if needThread {
		t, err := h.Agregator.GetThreadById(p.ThreadId)
		if err != nil {
			fmt.Println(err)
			c.JSON(404, err)
			return
		}
		response.Thread = &t
	}
	fmt.Println(response.Post.Created)
	c.JSON(200, response)
}
