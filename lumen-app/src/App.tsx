import React, { useState } from 'react';
import {connect, sendMenssage} from './api/send-service'
import Header from './components/header/header'

function App() {
  const [chatHistory, setChatHistory] = useState<any[]>([]);
  React.useEffect(() => {
    connect((msg) => {
      console.log("New Message");
      setChatHistory((prev) => [...prev, msg]);
    })
  }, []);

  function send():void{
    console.log("Mensagem enviada!")
    sendMenssage("Olá, teste")
  }

    return (
      <div className="App">
        <Header />
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
