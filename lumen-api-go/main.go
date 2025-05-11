package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler que realiza o upgrade da conexão HTTP para WebSocket e inicia o leitor de mensagens.
func websocketHandler(ctx *gin.Context) {
	// Realiza o upgrade da conexão HTTP para WebSocket
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Println(err)
	}
	// Chamada do leitor de mensagens
	reader(ws)
}

// Configura rota para o handler websocketHandler
func setupRoutes(router *gin.Engine) {
	router.GET("/lumen-chat/web-socket", websocketHandler)
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
