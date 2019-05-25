package handlers

import (
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) GetForumUsers(c *gin.Context) {
	slug := c.Param("slug")
	lim, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		lim = 100
	}

	since := c.Query("since")

	desc, _ := strconv.ParseBool(c.Query("desc"))

	exits, err := h.Agregator.IsForumExist(slug)
	if err != nil {
		c.JSON(404, err)
		return
	}

	if !exits {
		c.JSON(404, err)
		return
	}
	users := &[]models.User{}

	if desc {
		users, err = h.Agregator.GetUsersDESC(slug, lim, since)
	} else {
		users, err = h.Agregator.GetUsersASC(slug, lim, since)
	}
	if err != nil {
		c.JSON(404, err)
		return
	}

	if len(*users) == 0 {
		emptyArray := make([]int64, 0)
		c.JSON(200, emptyArray)
		return
	}

	c.JSON(200, users)
	return
}
