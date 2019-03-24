package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Floaty Square
type bestScore struct {
	name  string
	score int
}

func floatySquareHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxionlaptop/Documents/Code/Go/GameServer/html/floatysquare.html")
}

func floatySquareEchoHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		} else if string(msg) == "SOCKET_OPEN" {
			sortedScores := readTopScores()
			strScores := ""
			for i := 0; i < len(sortedScores); i++ {
				strScores += sortedScores[i].name + "\t" + strconv.Itoa(sortedScores[i].score) + "\n"
			}
			conn.WriteMessage(msgType, []byte(strScores))
		} else if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		} else {
			fmt.Println("Floaty Square New Score: " + string(msg))
			f, err := os.OpenFile("/home/haxxionlaptop/Documents/Code/Go/GameServer/txt/floatySquareScore.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			if _, err := f.Write([]byte(string(msg) + "\n")); err != nil {
				log.Fatal(err)
			}
			if err := f.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func readTopScores() []bestScore {
	file, err := os.Open("/home/haxxionlaptop/Documents/Code/Go/GameServer/txt/floatySquareScore.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fullList []bestScore
	var sortedList []bestScore
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		name := s[0]
		score, _ := strconv.Atoi(s[1])
		newScore := bestScore{name, score}
		fullList = append(fullList, newScore)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		highestScore := bestScore{"", 0}
		index := 0
		for j := 0; j < len(fullList); j++ {
			if fullList[j].score > highestScore.score {
				highestScore.name = fullList[j].name
				highestScore.score = fullList[j].score
				index = j
			}
		}
		sortedList = append(sortedList, highestScore)
		fullList = remove(fullList, index)
	}
	return sortedList
}

func remove(s []bestScore, i int) []bestScore {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
