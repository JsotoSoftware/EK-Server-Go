package network

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Allow all origins in development
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

// List of clients
var clients = make(map[*Client]bool)

func HandleConnections(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Fatal("Error upgrading to websocket: ", err)
	}
	defer conn.Close()

	// Register client
	client := &Client{Conn: conn, Send: make(chan []byte)}
	clients[client] = true

	log.Println("New client connected")

	// Read messages from client
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message: ", err)
			delete(clients, client)
			break
		}

		log.Println("Received message: ", string(msg))

		// Broadcast message to all clients
		for client := range clients {
			client.Send <- msg
		}
	}
}
