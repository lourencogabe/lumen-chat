import React, { useState } from 'react';
import {connect, sendMenssage} from './api/send-service'
import ChatInput from './components/ChatImput/ChatInput';
import Header from './components/header/Header';

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
      
    </div>
  )
}

export default App;
