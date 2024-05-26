package main

import (
	"fmt"
	"net/http"

	"go-socket-messaging-backend/pkg/websocket"
)

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket endpoint reached !")

	// this supposedly upgrades an http connection to
	// a websocket connection
	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	// we create a websocket client and register it to the pool
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}
	// since the pool has started on a goroutine
	// we send any info to it by means of channels
	pool.Register <- client

	// means by which client can read / receive messages
	client.Read()
}

func setupRoutes() {
	// before starting up the http server, we create a websocket pool
	// where multiple clients will use from the pool of connections.
	// NEED LITTLE BIT MORE DETAIL TBH
	pool := websocket.NewPool()
	// start the pool by means of a goroutine.
	// this handles the register, unregister and sending of messages to clients
	go pool.Start()

	// setting up the websocket http endpoint
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main() {
	// starting point of the server
	fmt.Println("starting server !")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
