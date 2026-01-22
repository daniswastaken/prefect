package main

import (
	"log"
	"net/http"
	"prefect/api"
)

func main() {
	// 1. The WebSocket API route
	// This maps the TypeScript "new WebSocket('ws://.../ws')" to your Go handler
	http.HandleFunc("/ws", api.StreamStats)

	// 2. The Static Web UI
	// This tells Go to serve index.html and styles.css from the /ui folder
	fs := http.FileServer(http.Dir("./ui"))
	http.Handle("/", fs)

	log.Println("Server starting on http://localhost:8080")

	// 3. Start the server
	// Using ":8080" allows it to be seen from Windows (WSL bridge)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
