package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	// creating a new pool instance
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	// starting the pool

	// infinite loop
	for {
		select {

		// to handle joining of client to the pool
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("size of connection pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				// sending message to all clients in the pool
				// that a new user has joined the chat
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined !"})
			}

		// to handle disconnection of client from the pool
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("size of connection pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				// sending message to all clients in the pool
				// that a user has disconnected from the chat
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected"})
			}

		// to handle the sending of messages from any client in the chat
		case message := <-pool.Broadcast:

			fmt.Println("Sending message to all clients in the pool")
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
