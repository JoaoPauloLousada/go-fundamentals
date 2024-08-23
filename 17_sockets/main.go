package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Init")
	m := NewManager()
	setupRoutes(m)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func wsHandler(m *Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrader error:", err)
		}
		log.Println("Client connected")
		err = ws.WriteMessage(1, []byte("Hello, Client!"))
		if err != nil {
			log.Println("writemessage error:", err)
		}

		m.Read(ws)
	}
}

func setupRoutes(m *Manager) {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler(m))
}
