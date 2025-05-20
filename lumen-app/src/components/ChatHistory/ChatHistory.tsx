import React from 'react';
import Message from '../Message/Message'; // ajuste o caminho conforme necessário
import './chat-history.scss';

// Tipo de uma mensagem individual
interface MessageType {
  data: string;
}

// Props do componente ChatHistory
interface ChatHistoryProps {
  chatHistory: MessageType[];
}

function ChatHistory({ chatHistory }: ChatHistoryProps) {
  return (
    <div className="ChatHistory">
      <h2>Chat History</h2>
      {chatHistory.map((msg, index) => (
        <Message key={index} message={msg.data} />
      ))}
    </div>
  );
}

export default ChatHistory;