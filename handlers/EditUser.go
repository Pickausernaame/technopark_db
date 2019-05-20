package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Pickausernaame/technopark_db/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

func (h *Handler) EditUser(c *gin.Context) {

	var editUser models.User
	_ = json.NewDecoder(c.Request.Body).Decode(&editUser)
	nickname := c.Param("nickname")
	curUser, err := h.Agregator.EditUserAgr(nickname, editUser)
	fmt.Println(err)
	if err != nil {
		if err.Error() == "no rows in result set" {
			c.JSON(404, e)
		} else if err.(pgx.PgError).Code == "23505" {
			c.JSON(409, e)
		}
		return
	}
	c.Writer.WriteHeader(200)
	_ = json.NewEncoder(c.Writer).Encode(curUser)
}
