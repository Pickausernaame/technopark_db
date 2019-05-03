package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"technopark_db/models"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var newUser models.User
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
			//err = json.NewEncoder(c.Writer).Encode(resUsers)
			return
		}
	}
	c.Writer.WriteHeader(201)

	//err = json.NewEncoder(c.Writer).Encode(newUser)
	//res, err := json.Marshal(newUser)
	c.JSON(201, newUser)
	if err != nil {
		fmt.Println("Error in CreateUser handler 3")
		fmt.Println(err)
	}
}
