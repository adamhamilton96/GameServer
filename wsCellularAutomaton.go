package main

import "net/http"

func gameOfLifeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/GameServer/html/gameoflife.html")
}

func langtonsAntHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/GameServer/html/langtonsant.html")
}

func briansBrainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/GameServer/html/briansbrain.html")
}

func wireWorldHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/GameServer/html/wireworld.html")
}

func sandPilesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/GameServer/html/sandpiles.html")
}

func cameraStipplingHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/GameServer/html/camera.html")
}
