package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) GetThreadDetails(c *gin.Context) {
	slug_or_id := c.Param("slug_or_id")
	id, err := strconv.Atoi(slug_or_id)
	if err != nil {
		fmt.Println(slug_or_id)
		// это slug
	} else {
		fmt.Println(id)
		// это id
	}
}
