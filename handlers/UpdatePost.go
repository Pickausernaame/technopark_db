package handlers

import (
	"encoding/json"
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, err)
		return
	}

	var req models.Post

	//buf, _ := c.GetRawData()
	//err = req.UnmarshalJSON(buf)
	_ = json.NewDecoder(c.Request.Body).Decode(&req)
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
