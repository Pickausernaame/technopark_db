package handlers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"technopark_db/models"
)

func (h *Handler) GetThreadDetails(c *gin.Context) {
	slug_or_id := c.Param("slug_or_id")
	id, err := strconv.Atoi(slug_or_id)
	var thread models.Thread
	if err != nil {
		thread, err = h.Agregator.GetThreadAgr(slug_or_id)
		if err != nil {
			c.JSON(404, err)
			return
		}
	} else {
		thread, err = h.Agregator.GetThreadById(id)
		if err != nil {
			c.JSON(404, err)
			return
		}
	}
	c.JSON(200, thread)

}
