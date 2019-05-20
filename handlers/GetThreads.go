package handlers

import (
	"encoding/json"
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func (h *Handler) GetThreads(c *gin.Context) {
	var input []models.Thread
	_ = json.NewDecoder(c.Request.Body).Decode(&input)
	slug := c.Param("slug")
	limitStr, _ := c.GetQuery("limit")
	limit, _ := strconv.Atoi(limitStr)
	sinceStr, _ := c.GetQuery("since")
	descStr, _ := c.GetQuery("desc")
	desc, _ := strconv.ParseBool(descStr)
	since, err := time.Parse(time.RFC3339, sinceStr)
	if err != nil {
		if desc {
			since = time.Unix(64060588800, 0)
		} else {
			since = time.Unix(0, 0)
		}
	}
	threads, err, exist := h.Agregator.GetThreadsAgr(slug, limit, since, desc)
	if err != nil || !exist {
		c.JSON(404, e)
		return
	}
	if threads == nil {
		emptyArray := make([]int64, 0)
		c.JSON(200, emptyArray)
		return
	}
	c.JSON(200, threads)
	return
}
