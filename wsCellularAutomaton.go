package main

import "net/http"

func gameOfLifeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/Documents/Programming/GameServer/html/gameoflife.html")
}

func langtonsAntHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/Documents/Programming/GameServer/html/langtonsant.html")
}

func briansBrainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/Documents/Programming/GameServer/html/briansbrain.html")
}

func wireWorldHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/Documents/Programming/GameServer/html/wireworld.html")
}

func sandPilesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/Documents/Programming/GameServer/html/sandpiles.html")
}

func cameraStipplingHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/Documents/Programming/GameServer/html/camera.html")
}
