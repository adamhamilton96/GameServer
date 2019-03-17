package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxionlaptop/Documents/Code/Go/GameServer/html/home.html")
}

func main() {
	fmt.Println("Game Server Online")
	fmt.Print(boardString())
	// Home
	http.HandleFunc("/", rootHandler)

	// Connect4
	http.HandleFunc("/connect4/", connect4Handler)
	http.HandleFunc("/connect4echo", connect4EchoHandler)

	log.Fatal(http.ListenAndServe("192.168.1.230:1255", nil))
}
