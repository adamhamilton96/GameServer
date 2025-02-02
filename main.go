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
	http.ServeFile(w, r, "/home/haxxion/GameServer/html/home.html")
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
	
	fs := http.FileServer(http.Dir("/home/haxxion/GameServer/"))
	http.Handle("/benn/", addHeaders(fs))
	http.Handle("/pedro.mp4", addVideoHeaders(fs))
	http.Handle("/pedro.mp3", addSoundHeaders(fs))

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

// Benn
func addHeaders(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Cross-Origin-Resource-Policy", "cross-origin")
		fs.ServeHTTP(w, r)
	}
}

func addVideoHeaders(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Cross-Origin-Resource-Policy", "cross-origin")
		w.Header().Set("Content-Type", "video/mp4")
		fs.ServeHTTP(w, r)
	}
}

func addSoundHeaders(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Cross-Origin-Resource-Policy", "cross-origin")
		w.Header().Set("Content-Type", "audio/mp3")
		fs.ServeHTTP(w, r)
	}
}
