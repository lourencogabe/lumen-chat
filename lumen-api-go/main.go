package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pkg "github.com/lourencogabe/lumen-chat/pkg/websocket"
)

// Configura rota para o handler websocketHandler
func setupRoutes(router *gin.Engine) {
	router.GET("/lumen-chat/web-socket", pkg.WebSocketHandler)
}

func main() {
	route := gin.Default()

	route.GET("/lumen-chat", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Lumen Chat v0.01",
		})
	})

	setupRoutes(route)

	route.Run(":7070")
}
