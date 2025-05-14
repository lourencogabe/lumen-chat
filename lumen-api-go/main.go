package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lourencogabe/lumen-chat/pkg/websocket"
)

// A função serveWS() configura o endpoint para o websocket
func serveWs(pool *websocket.Pool, ctx *gin.Context) {
	fmt.Println("WebSocket Endpoint Hit")
	// Configura conn para chamar WebSocketHandler e atualizar websocket com o client.
	conn, err := websocket.WebSocketHandler(ctx)
	if err != nil {
		log.Println(err)
	}
	// Inicializa a struct Client.
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	// Registra o client no pool e chama a função Listen para controle.
	pool.Register <- client
	client.Listen()
}

// Configura rota para o handler websocketHandler
func setupRoutes(router *gin.Engine) {
	// Inicia um novo pool
	pool := websocket.NewPool()
	// Inicia goruntime
	go pool.Start()
	// Configura rota
	router.GET("/lumen-chat/web-socket", func(ctx *gin.Context) {
		serveWs(pool, ctx)
	})
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
