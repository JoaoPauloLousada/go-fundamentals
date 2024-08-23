package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Read(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("error:", err)
			return
		}

		log.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("conn write message error:", err)
			return
		}
	}
}
