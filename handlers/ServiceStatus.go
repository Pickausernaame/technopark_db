package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ServiceStatus(c *gin.Context) {
	status, err := h.Agregator.Status()
	if err != nil {
		fmt.Println(err)
		c.JSON(404, status)
		return
	}
	c.JSON(200, status)
}
