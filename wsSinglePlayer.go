package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// Floaty Square
func floatySquareHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxionlaptop/Documents/Code/Go/GameServer/html/floatysquare.html")
}

func floatySquareEchoHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		} else if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		} else {
			fmt.Println(msg)
			num, _ := strconv.Atoi(string(msg))
			fmt.Println(num)
			f, err := os.Create("/home/haxxionlaptop/Documents/Code/Go/GameServer/txt/f,floatySquareScores")
			check(err)
			defer f.Close()

			n3, err := f.WriteString(strconv.Itoa(num) + "\n")
			fmt.Printf("wrote %d bytes\n", n3)
			f.Sync()
			fmt.Println("written")
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
