package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateThread(c *gin.Context) {
	fmt.Println("Create thread" + c.Param("slug_or_id"))
}
