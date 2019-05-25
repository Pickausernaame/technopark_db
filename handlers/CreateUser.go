package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

func (h *Handler) CreateUser(c *gin.Context) {
	newUser := &models.User{}
	err := json.NewDecoder(c.Request.Body).Decode(&newUser)
	if err != nil {
		fmt.Println("Error in CreateUser handler 3")
		fmt.Println(err)
	}
	newUser.Nickname = c.Param("nickname")
	err = h.Agregator.CreateUserAgr(newUser)
	if err != nil {
		if err.(pgx.PgError).Code == "23505" {
			resUsers, err := h.Agregator.ErrorCreateUserArg(newUser.Nickname, newUser.Email)
			if err != nil {
				fmt.Println("Error in CreateUser handler 2")
				fmt.Println(err)
			}
			c.Writer.WriteHeader(409)
			c.JSON(409, resUsers)
			return
		}
	}
	c.JSON(201, newUser)
}
