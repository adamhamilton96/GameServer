package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Floaty Square
type bestScore struct {
	name  string
	score int
}

func floatySquareHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/Documents/Programming/GameServer/html/floatysquare.html")
}

func floatySquareEchoHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		} else if string(msg) == "SOCKET_OPEN" {
			fmt.Println("User playing Floaty Square")
			sortedScores := readTopScores("floatySquareScore.txt")
			strScores := ""
			for i := 0; i < len(sortedScores); i++ {
				strScores += sortedScores[i].name + "\t" + strconv.Itoa(sortedScores[i].score) + "\n"
			}
			conn.WriteMessage(msgType, []byte(strScores))
		} else if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		} else {
			fmt.Println("Floaty Square New Score: " + string(msg))
			if checkString(string(msg)) {
				f, err := os.OpenFile("/home/haxxion/Documents/Programming/GameServer/txt/floatySquareScore.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Fatal(err)
				}
				if _, err := f.Write([]byte(string(msg) + "\n")); err != nil {
					log.Fatal(err)
				}
				if err := f.Close(); err != nil {
					log.Fatal(err)
				}
			} else {
				fmt.Println("Cheater detected")
				fmt.Println(string(msg) + " not accepted")
				conn.WriteMessage(msgType, []byte("Stop cheating loser"))
			}
		}
	}
}

// Snake
func snakeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/haxxion/Documents/Programming/GameServer/html/snake.html")
}

func snakeEchoHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		} else if string(msg) == "SOCKET_OPEN" {
			fmt.Println("User playing Snake")
			sortedScores := readTopScores("snakeScore.txt")
			strScores := ""
			for i := 0; i < len(sortedScores); i++ {
				strScores += sortedScores[i].name + "\t" + strconv.Itoa(sortedScores[i].score) + "\n"
			}
			conn.WriteMessage(msgType, []byte(strScores))
		} else if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		} else {
			fmt.Println("Snake New Score: " + string(msg))
			if checkString(string(msg)) {
				f, err := os.OpenFile("/home/haxxion/Documents/Programming/GameServer/txt/snakeScore.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Fatal(err)
				}
				if _, err := f.Write([]byte(string(msg) + "\n")); err != nil {
					log.Fatal(err)
				}
				if err := f.Close(); err != nil {
					log.Fatal(err)
				}
			} else {
				fmt.Println("Cheater detected")
				fmt.Println(string(msg) + " not accepted")
				conn.WriteMessage(msgType, []byte("Stop cheating loser"))
			}
		}
	}
}

func checkString(msg string) bool {
	words := strings.Fields(msg)
	if len(words) != 2 {
		return false
	}

	if len(words[0]) > 20 {
		return false
	}

	if !strings.ContainsAny(words[0], "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	}

	if !strings.ContainsAny(words[1], "0123456789") {
		return false
	}

	for _, r := range msg {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsSpace(r) {
			return false
		}
	}
	score, err := strconv.Atoi(words[1])
	if err != nil {
		return false
	}

	if strings.ToLower(words[0]) == "Tom" && (score < 1 || score > 100) {
		return false
	}
	if strings.ToLower(words[0]) == "JAMES" && (score < 1 || score > 20) {
		return false
	}
	if score < 1 || score > 500 {
		return false
	}

	return true
}

func readTopScores(path string) []bestScore {

	// Open file
	file, err := os.Open("/home/haxxion/Documents/Programming/GameServer/txt/" + path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fullList []bestScore
	var sortedList []bestScore
	// for each line in txt file
	for scanner.Scan() {
		// split into name and score
		s := strings.Split(scanner.Text(), " ")
		name := s[0]
		score, _ := strconv.Atoi(s[1])
		newScore := bestScore{name, score}
		// add to array
		fullList = append(fullList, newScore)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// top 10 best scores
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
		// add to top 10 list
		sortedList = append(sortedList, highestScore)
		// remove original so not refound
		fullList = remove(fullList, index)
	}
	return sortedList
}

func remove(s []bestScore, i int) []bestScore {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
