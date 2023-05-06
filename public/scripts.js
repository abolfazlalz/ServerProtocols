const sendBtn = document.getElementById("btn-send")
const messageInput = document.getElementById("input-message")

const ws = new WebSocket("ws://localhost:8000/ws")

function handleSendClick() {
    const message = messageInput.value
    messageInput.value = ""
    ws.send(message)
}

sendBtn.addEventListener('click', handleSendClick)
