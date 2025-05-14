import React, { useState } from 'react';
import {connect, sendMenssage} from './api/send-service'
import Header from './components/header/header'
import ChatInput from './components/chat-imput/chat-input';

function App() {
  const [chatHistory, setChatHistory] = useState<any[]>([]);
  React.useEffect(() => {
    connect((msg) => {
      console.log("New Message");
      setChatHistory((prev) => [...prev, msg]);
    })
  }, []);

  function send(event: React.KeyboardEvent<HTMLInputElement>){
    if(event.key === 'Enter'){
      sendMenssage(event.currentTarget.value)
      event.currentTarget.value = ''
    }

    console.log("Mensagem enviada!")
    sendMenssage("Olá, teste")
  }

    return (
      <div className="App">
        <Header />
        <ChatInput send={send}/>
      <div>
        {chatHistory.map((msg, index) => (
          <div key={index}>
            {typeof msg.data === "string" ? msg.data : JSON.stringify(msg.data)}
          </div>
        ))}
      </div>
      <button onClick={send}>Hit</button>
    </div>
  )
}

export default App;
