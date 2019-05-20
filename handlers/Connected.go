package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) Connected(c *gin.Context) {

	c.JSON(201, "connected")
}
