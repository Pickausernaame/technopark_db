package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetForum(c *gin.Context) {
	forum, err := h.Agregator.GetForumAgr(c.Param("slug"))
	if err != nil {
		c.JSON(404, e)
		return
	}
	c.JSON(200, forum)
}
