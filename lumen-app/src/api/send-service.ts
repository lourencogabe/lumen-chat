import ChatHystory from "../components/ChatHistory/ChatHistory"

const socket = new WebSocket("ws://localhost:7070/lumen-chat/web-socket")

function connect(cb: (msg: MessageEvent) => void): void{
    console.log("Aguardando conexão...")

    socket.onopen = () => { 
        console.log("Conexão realizada com sucesso!")
    }

    socket.onmessage = msg => {
        cb(msg)
    }

    socket.onclose = event => {
        console.log("Porta de conexão fechada: ",event)
    }

    socket.onerror = error => {
        console.log("Socket Error: ", error);
      };
}

function sendMenssage(msg:string): string {
    console.log("Enviando mensagem: ", msg);
    socket.send(msg);
    return ''
}

export{ connect, sendMenssage }