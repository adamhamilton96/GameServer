package main

import (
	"crypto/tls"
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

var (
	CertFilePath = "/etc/letsencrypt/live/haxxion.hopto.org/fullchain.pem"
	KeyFilePath  = "/etc/letsencrypt/live/haxxion.hopto.org/privkey.pem"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/Documents/Programming/GameServer/html/home.html")
}

func main() {
	// load tls certificates
	serverTLSCert, err := tls.LoadX509KeyPair(CertFilePath, KeyFilePath)
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
	}
	fmt.Println("haxxion.hopto.org Online :)")

	// Home
	http.HandleFunc("/", rootHandler)

	// Multiplayer
	// Connect4
	http.HandleFunc("/connect4/", connect4Handler)
	http.HandleFunc("/connect4echo", connect4EchoHandler)

	// Singleplayer
	// FloatySquare
	http.HandleFunc("/floatysquare/", floatySquareHandler)
	http.HandleFunc("/floatysquareecho", floatySquareEchoHandler)
	// Snake
	http.HandleFunc("/snake/", snakeHandler)
	http.HandleFunc("/snakeecho", snakeEchoHandler)

	// Cellular Automata
	http.HandleFunc("/gameoflife/", gameOfLifeHandler)
	http.HandleFunc("/langtonsant/", langtonsAntHandler)
	http.HandleFunc("/briansbrain/", briansBrainHandler)
	http.HandleFunc("/wireworld/", wireWorldHandler)
	http.HandleFunc("/sandpiles/", sandPilesHandler)

	// Camera
	http.HandleFunc("/stippling/", cameraStipplingHandler)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverTLSCert},
	}
	server := http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
	}
	defer server.Close()
	log.Fatal(server.ListenAndServeTLS("", ""))
}
