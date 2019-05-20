package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"technopark_db/models"
)

func (h *Handler) UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, err)
		return
	}

	var req models.Post
	buf, _ := c.GetRawData()
	err = req.UnmarshalJSON(buf)
	if err != nil {
		c.JSON(404, err)
		return
	}

	post, err := h.Agregator.GetPost(id)

	if err != nil {
		c.JSON(404, err)
		return
	}

	if req.Message == "" || req.Message == post.Message {
		c.JSON(200, post)
		return
	}

	err = h.Agregator.PostUpdate(id, req.Message)
	if err != nil {
		c.JSON(404, err)
	}

	post.Message = req.Message
	post.IsEdited = true

	c.JSON(200, post)
	return
}
