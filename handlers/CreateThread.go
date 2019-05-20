package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

func (h *Handler) CreateThread(c *gin.Context) {
	var thread models.Thread
	_ = json.NewDecoder(c.Request.Body).Decode(&thread)
	thread.Forum = c.Param("slug")
	resThread, err := h.Agregator.CreateThreadAgr(thread)
	if err != nil {
		if err.(pgx.PgError).Code == "23505" {
			resThread, err = h.Agregator.GetThreadAgr(thread.Slug)
			if err != nil {
				fmt.Println(err)
			}
			c.JSON(409, resThread)
			return
		} else if err.(pgx.PgError).Code == "23502" {
			c.JSON(404, e)
			return
		}
	}
	c.JSON(201, resThread)
}
