// Create WebSocket connection.
const socket = new WebSocket("ws://localhost:8080");

// Connection opened
socket.addEventListener("open", (event) => {
    console.log("Connection opened.")
});

// Listen for messages
socket.addEventListener("message", (event) => {
  document.getElementById('response').innerHTML = event.data;
});

function send() {
	var command = document.getElementById('command').value;
    // console.log(command)
    socket.send(command)
}

function commandKeyDown(ele) {
	if (event.key === 'Enter') {
		send();
	}
}
