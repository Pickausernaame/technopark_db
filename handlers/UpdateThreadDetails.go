package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) UpdateThreadDetails(c *gin.Context) {
	var update models.Thread
	err := json.NewDecoder(c.Request.Body).Decode(&update)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, err)
	}
	slug_or_id := c.Param("slug_or_id")
	id, err := strconv.Atoi(slug_or_id)
	response := &models.Thread{}
	if err != nil {
		response, err = h.Agregator.GetThreadAgr(slug_or_id)
		if err != nil {
			c.JSON(404, err)
			return
		}
	} else {
		response, err = h.Agregator.GetThreadById(id)
		if err != nil {
			c.JSON(404, err)
			return
		}
	}
	if update.Message == "" && update.Title == "" {
		c.JSON(200, response)
		return
	}

	if update.Message == "" {
		update.Message = response.Message
	}

	if update.Title == "" {
		update.Title = response.Title
	}
	update.Id = response.Id
	response.Message = update.Message
	response.Title = update.Title
	err = h.Agregator.ThreadUpdate(&update)
	if err != nil {
		c.JSON(404, err)
		return
	}
	c.JSON(200, response)
}
