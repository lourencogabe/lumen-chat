package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	//Permite requisições de qualquer origem
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	route := gin.Default()

	route.GET("/lumen-chat/get/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Lumen Chat v0.01",
		})
	})

	route.Run(":7070")
}
