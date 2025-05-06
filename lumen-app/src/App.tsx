import React from 'react';
import logo from './logo.svg';
import {connect, sendMenssage} from './api/send-service'

function App() {
  connect()

  function send(): any{
    console.log("hello");
    sendMenssage("hello");
  }

  return (
    <div className="App">
    <button onClick={send}>Hit</button>
  </div>
  );
}

export default App;
