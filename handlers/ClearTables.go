package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) ClearTables(c *gin.Context) {
	h.Agregator.ClearTableAgr()
}
