import React from 'react';
import {connect, sendMenssage} from './api/send-service'
import Header from './components/header/header'

function App() {
  React.useEffect(() => {
    connect();
  }, []);

  function send(): any{
    console.log("hello");
    sendMenssage("hello");
  } 
    return (
      <div className="App">
        <Header />
        <button onClick={send}>Hit</button>
      </div>
    );
}

export default App;
