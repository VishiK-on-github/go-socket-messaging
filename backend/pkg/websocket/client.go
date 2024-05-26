package websocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {

	// delegate closing of connection / client disconnection
	// after the read method completes
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {

		// get the message
		messageType, p, err := c.Conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		// create a message instance
		message := Message{Type: messageType, Body: string(p)}

		// send the message over the broadcast channel
		c.Pool.Broadcast <- message

		fmt.Printf("Message received: %+v\n", message)
	}
}
