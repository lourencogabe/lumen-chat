package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Define o tamanho do buffer para leitura e escrita.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// Permite requisições de qualquer origem.
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Le a mensagem recebida.
func reader(conn *websocket.Conn) {
	for {
		//Le a mensagem e armazena no payload (p).
		messagetype, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))
		// Reenvia a mesma mensagem para confirmar comunicação servidor-cliente.
		if err := conn.WriteMessage(messagetype, p); err != nil {
			log.Println(err)
			return
		}
	}
}

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
