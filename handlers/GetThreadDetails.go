package handlers

import (
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) GetThreadDetails(c *gin.Context) {
	slug_or_id := c.Param("slug_or_id")
	id, err := strconv.Atoi(slug_or_id)
	thread := &models.Thread{}
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
