const themeList = document.getElementById('theme-list');
const chatMessages = document.getElementById('chat-messages');
const messageInput = document.getElementById('message-input');
const sendBtn = document.getElementById('send-btn');

fetch('http://localhost:8080/themes')
    .then(response => response.json())
    .then(data => {
        const themes = data.themes;
        themes.forEach(theme => {
            const li = document.createElement('li');
            li.textContent = theme;
            themeList.appendChild(li);
        });
    })
    .catch(err => console.log(err));

const socket = new WebSocket('ws://localhost:8080/ws');

socket.onmessage = function (event) {
    const message = event.data;
    const p = document.createElement('p');
    p.textContent = message;
    chatMessages.appendChild(p);
    chatMessages.scrollTop = chatMessages.scrollHeight; // Auto-scroll to the bottom
};

sendBtn.addEventListener('click', () => {
    const message = messageInput.value.trim();
    if (message) {
        socket.send(message);
        messageInput.value = '';
    }
});

messageInput.addEventListener('keydown', (e) => {
    if (e.key === 'Enter') {
        sendBtn.click();
    }
});
