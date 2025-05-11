package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"

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
func ReaderMsg(conn *websocket.Conn) {
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

func WriterMsg(conn *websocket.Conn) {
	for {
		fmt.Println("Sending")
		messageType, reader, err := conn.NextReader()
		if err != nil {
			log.Println(err)
			return
		}
		write, err := conn.NextWriter(messageType)
		if err != nil {
			log.Println(err)
			return
		}
		if _, err := io.Copy(write, reader); err != nil {
			log.Println(err)
			return
		}
	}

}
