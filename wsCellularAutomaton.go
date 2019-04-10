package main

import "net/http"

func gameOfLifeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/mnt/e/Programming/Go/GameServer/html/gameoflife.html")
}

func langtonsAntHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/mnt/e/Programming/Go/GameServer/html/langtonsant.html")
}

func briansBrainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/mnt/e/Programming/Go/GameServer/html/briansbrain.html")
}

func wireWorldHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/mnt/e/Programming/Go/GameServer/html/wireworld.html")
}

func sandPilesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/mnt/e/Programming/Go/GameServer/html/sandpiles.html")
}
