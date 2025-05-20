import React from 'react';
import './Message.scss';

type MessageProps = {
  message: string; // JSON string
};

function Message({ message }: MessageProps) {
  const parsed = JSON.parse(message);
  return <div className="Message">{parsed.body}</div>;
}

export default Message;