package services

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	clients      = make(map[*websocket.Conn]bool)
	clientsMutex sync.Mutex
	upgrader     = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
)

// BroadcastCounters sends counters to all connected WebSocket clients
func BroadcastCounters(counters interface{}) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
// 	counters, err := fetchAllCounters()
// 	if err != nil {
// 		log.Println("Error fetching counters for broadcast:", err)
// 		return
// 	}
	for client := range clients {
		err := client.WriteJSON(counters)
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}

// HandleWebSocket handles new WebSocket connections
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("WebSocket connection error:", err)
        return
    }

    clientsMutex.Lock()
    clients[conn] = true
    clientsMutex.Unlock()

    // Clean up on close
    defer func() {
        clientsMutex.Lock()
        delete(clients, conn)
        clientsMutex.Unlock()
        conn.Close()
    }()

    // Listen for messages (if needed)
    for {
        _, _, err := conn.ReadMessage()
        if err != nil {
            log.Println("WebSocket read error:", err)
            break
        }
    }
}
