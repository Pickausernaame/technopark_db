package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Connected(c *gin.Context) {
	//err := h.Agregator.Clastering()
	//fmt.Println(err)
	//if err != nil {
	//	c.JSON(409, "problems")
	//}
	c.JSON(201, "connected")
}
