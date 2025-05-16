import { useState } from "react"
import "./ChatInput.scss"

function ChatInput(e: any){
    const [valor, setValor] =useState('')

    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        setValor(e.target.value);
    }

    return (
        <div className="ChatInput">
            <input type="text" value={valor} onChange={handleChange} />
        </div>
    )
}

export default ChatInput