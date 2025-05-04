package main

import (
	"log"
	"net/http"
	"network-monitor/ws"
)

func main() {
	http.HandleFunc("/ws", ws.HandleWS)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
