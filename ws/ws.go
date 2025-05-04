package ws

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type ServerStatus struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    string `json:"port"`
	Status  string `json:"status"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWS(w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket request received")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	servers := []ServerStatus{
		{"Google", "google.com", "80", ""},
		{"GitHub", "github.com", "443", ""},
		{"InvalidHost", "invalid.host", "80", ""},
	}

	for {
		for i := range servers {
			servers[i].Status = checkServer(servers[i].Address, servers[i].Port)
		}

		err := conn.WriteJSON(servers)
		if err != nil {
			log.Println("Write error:", err)
			break
		}

		time.Sleep(5 * time.Second)
	}
}

func checkServer(address, port string) string {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(address, port), 2*time.Second)
	if err != nil {
		return "offline"
	}
	conn.Close()
	return "online"
}
