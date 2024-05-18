var socket = new WebSocket("ws://localhost:9000/ws");

let connect = (cb) => {
	console.log("connecting...");

	socket.onopen = () => {
		console.log("connected to socket !");
	};

	socket.onmessage = (msg) => {
		console.log("message from websocket: ", msg);
	};

	socket.onclose = (e) => {
		console.log("socket closed connection: ", e);
	};

	socket.onerror = (err) => {
		console.log("socket error: ", err);
	};
};

let sendMsg = (msg) => {
	console.log("sending message: ", msg);
	socket.send(msg);
};

export { connect, sendMsg };
