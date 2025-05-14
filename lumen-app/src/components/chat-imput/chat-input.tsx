import { useState } from "react"
import "./chat-input.scss"

function ChatInput(e: any){
    const [valor, setValor] =useState('')

    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        setValor(e.target.value);
    }

    return (
        <div>
            <input type="text" value={valor} onChange={handleChange} />
        </div>
    )
}

export default ChatInput