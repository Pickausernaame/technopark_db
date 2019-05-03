package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"technopark_db/models"
)

func (h *Handler) CreateForum(c *gin.Context) {
	if c.Param("slug") != "create" {
		fmt.Println(c.Param("slug"))
		return
	}
	var newForum models.Forum
	var resForum models.Forum
	err := json.NewDecoder(c.Request.Body).Decode(&newForum)
	if err != nil {
		fmt.Println("Error CreateForum Handler:")
		fmt.Println(err)
	}
	resForum, err = h.Agregator.CreateForumAgr(&newForum)

	if err != nil {
		fmt.Println("Error in Agregation:")
		fmt.Println(err)
		// такой форум есть
		if err.(pgx.PgError).Code == "23505" {
			//c.Writer.WriteHeader(409)
			resForum, err = h.Agregator.ErrorCreateForumAgr(newForum.Slug)
			fmt.Println("BAD SLUG")
			fmt.Println(resForum)
			c.JSON(409, resForum)
			//err = json.NewEncoder(c.Writer).Encode(resForum)
			return
			// такого пользователя нет
		} else if err.(pgx.PgError).Code == "23502" {
			c.JSON(404, e)
			return
		}

	}
	c.JSON(201, resForum)
	//c.Writer.WriteHeader(201)
	//err = json.NewEncoder(c.Writer).Encode(resForum)
	if err != nil {
		fmt.Println("Error in Encoding:")
		fmt.Println(err)
	}
}
