const socket = new WebSocket("ws://localhost:7070/lumen-chat/web-socket")

function connect(): any {
    console.log("Aguardando conexão...")

    socket.onopen = () => { 
        console.log("Conexão realizada com sucesso!")
    }

    socket.onmessage = msg => {
        console.log(msg)
    }

    socket.onclose = event => {
        console.log("Porta de conexão fechada: ",event)
    }

    socket.onerror = error => {
        console.log("Socket Error: ", error);
      };
}

function sendMenssage(msg): any {
    console.log("Enviando mensagem: ", msg);
    socket.send(msg);
}

export{ connect, sendMenssage }