package main

import "net/http"

func gameOfLifeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxionlaptop/Documents/Code/Go/GameServer/html/gameoflife.html")
}
