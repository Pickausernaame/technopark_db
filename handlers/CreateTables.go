package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) CreateTables(c *gin.Context) {
	h.Agregator.CreateTableAgr()
}
