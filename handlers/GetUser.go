package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(c *gin.Context) {
	curUser, err := h.Agregator.GetUserAgr(c.Param("nickname"))
	if err != nil {
		fmt.Println(err)
		c.JSON(404, e)
		return
	}
	//c.Writer.WriteHeader(200)
	c.JSON(200, curUser)
	//_ = json.NewEncoder(c.Writer).Encode(curUser)

}
