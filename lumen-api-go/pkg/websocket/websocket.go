package websocket

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

// A função Listen() escuta continuamente as mensagens recebidas de um cliente WebSocket.
// Quando uma mensagem chega, ela é empacotada em um struct Message e enviada para todos os clientes através do canal Broadcast.
func (c *Client) Listen() {
	// Adia a finalização da conexao e do worker até que o método Read finalize seu processamento.
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Println("Message Received: %+v\n", message)
	}
}

// Define o tamanho do buffer para leitura e escrita.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// Permite requisições de qualquer origem.
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Handler que realiza o upgrade da conexão HTTP para WebSocket.
func WebSocketHandler(ctx *gin.Context) (*websocket.Conn, error) {
	// Realiza o upgrade da conexão HTTP para WebSocket
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Println(err)
	}
	// Retorna a conexão
	return conn, nil
}

// Le a mensagem recebida.
func ReaderMsg(conn *websocket.Conn) {
	for {
		// Le a mensagem e armazena no payload (p).
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

// Reenvia mesma msg
func WriterMsg(conn *websocket.Conn) {
	for {
		fmt.Println("Sending")
		// Captura a proxima msg recebida.
		messageType, reader, err := conn.NextReader()
		if err != nil {
			log.Println(err)
			return
		}
		// Recebe a conexão para envio de uma nova msg.
		write, err := conn.NextWriter(messageType)
		if err != nil {
			log.Println(err)
			return
		}
		// Le todo o conteudo de reader e armazena em write para envio.
		if _, err := io.Copy(write, reader); err != nil {
			log.Println(err)
			return
		}
		// Fecha a conexão após cada envio.
		if err := write.Close(); err != nil {
			log.Println(err)
			return
		}
	}

}
