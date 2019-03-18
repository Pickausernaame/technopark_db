package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"models"
	"strconv"
)

func (h *Handler) CreatePost(c *gin.Context) {
	var input []models.Post
	json.NewDecoder(c.Request.Body).Decode(&input)
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
