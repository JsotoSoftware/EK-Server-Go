package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ek-server/internal/network"
)

func main() {
	// Websocket route
	http.HandleFunc("/ws", network.HandleConnections)

	// Start the server
	log.Println("Starting server on port 8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Websocket server started on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe failed: ", err)
	}
}
