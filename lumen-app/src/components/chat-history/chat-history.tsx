import "./chat-history.scss"

//tipo do dado
interface Message{
    data: string;
}

//vetor tipado
interface ChatHistory{
    chatHistory: Message[]
}


function ChatHystory({chatHistory}: ChatHistory) {
    return(
        <div className="ChatHistory">
            <h2>Chat History</h2>
                {/* realiza um map para pegar o primeiro valor do vetor e armazenar na propriedade msg*/}
                {chatHistory.map((msg, index) => (
                    //Pega o .data da linha em que o index se encontra
                    <p key={index}>{msg.data}</p>
                ))}
        </div>
    )
}

export default ChatHistory;