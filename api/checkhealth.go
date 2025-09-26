package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) checkHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"staus": "ok",
		"message": "server is healthy",
	})
}