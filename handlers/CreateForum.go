package handlers

import (
	"encoding/json"
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

var (
	BadUnmarshal models.Errors
	UserNotExist models.Errors
)

func init() {
	BadUnmarshal.Error = "Bad Unmarshal"
	UserNotExist.Error = "User is not exist"
}

func (h *Handler) CreateForum(c *gin.Context) {
	if c.Param("slug") != "create" {
		return
	}
	var newForum models.Forum
	var resForum models.Forum
	err := json.NewDecoder(c.Request.Body).Decode(&newForum)
	if err != nil {
		c.JSON(409, BadUnmarshal)
	}
	resForum, err = h.Agregator.CreateForumAgr(&newForum)
	if err != nil {
		// такой форум есть
		if err.(pgx.PgError).Code == "23505" {
			resForum, err = h.Agregator.ErrorCreateForumAgr(newForum.Slug)
			c.JSON(409, resForum)
			return
			// такого пользователя нет
		} else if err.(pgx.PgError).Code == "23502" {
			c.JSON(404, UserNotExist)
			return
		}
	}
	c.JSON(201, resForum)
}
