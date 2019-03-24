package main

import (
	"fmt"
	"log"
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
			num, _ := strconv.Atoi(string(msg))
			fmt.Println(num)

			f, err := os.OpenFile("/home/haxxionlaptop/Documents/Code/Go/GameServer/txt/floatySquareScore.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			if _, err := f.Write([]byte(strconv.Itoa(num) + "\n")); err != nil {
				log.Fatal(err)
			}
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
			f.Close()
		}
	}
}
