package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

type player struct {
	clientConn *websocket.Conn
}

// Board must be in form y,x
var players []player
var board [][]int
var boardHeight int
var boardLength int
var turn int
var winner int

func setup() {
	fmt.Println("Connect 4: Setting up board")
	turn = 1
	winner = 0
	boardHeight = 6
	boardLength = 7
	board = make([][]int, boardHeight)
	for j := range board {
		board[j] = make([]int, boardLength)
	}
	fmt.Println("Connect 4: Board setup complete")
}

func boardString() string {
	boardString := ""
	for j := 0; j < boardHeight; j++ {
		for i := 0; i < boardLength; i++ {
			boardString += strconv.Itoa(board[j][i]) + " "
		}
	}
	return boardString
}

func dropToken(i int) bool {
	fmt.Println(strconv.Itoa(i))
	placed := false
	j := boardHeight - 1
	for placed == false {
		if board[j][i] == 0 {
			board[j][i] = turn
			placed = true
			if checkWinner() == 1 {
				winner = 1
			} else if checkWinner() == 2 {
				winner = 2
			}
			if checkDraw() == true {
				winner = -1
			}
			if turn == 1 {
				turn = 2
				fmt.Println("Connect 4: now P2 turn")
			} else if turn == 2 {
				turn = 1
				fmt.Println("Connect 4: now P1 turn")
			}
			return true
		} else if j > 0 {
			j--
		} else {
			return false
		}
	}
	fmt.Println("Should never get here 61")
	return true
}

func connect4EchoHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if string(msg) == "SOCKET_OPEN" {
			if len(players) == 0 {
				setup()
			}
			newPlayer := player{conn}
			players = append(players, newPlayer)
			fmt.Println("Connect 4: Player has joined")
			fmt.Println("Connect4: Players len:" + strconv.Itoa(len(players)))
			if len(players) == 2 {
				fmt.Println("Connect 4: Game full")
				for i := 0; i < len(players); i++ {
					if i == turn-1 {
						players[i].clientConn.WriteMessage(msgType, []byte(boardString()+"1"))
					} else {
						players[i].clientConn.WriteMessage(msgType, []byte(boardString()+"0"))
					}
				}
			}
		} else if string(msg) == "SOCKET_CLOSED" {
			for i := 0; i < len(players); i++ {
				if players[i].clientConn.RemoteAddr() == conn.RemoteAddr() {
					fmt.Printf("%s removed from slice\n", players[i].clientConn.RemoteAddr())
					players = append(players[:i], players[i+1:]...)
					fmt.Println("players len:" + strconv.Itoa(len(players)))
				}
			}
		} else if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		} else {
			// Input = number
			num, _ := strconv.Atoi(string(msg))
			dropToken(num)
			for i := 0; i < len(players); i++ {
				if i == turn-1 {
					players[i].clientConn.WriteMessage(msgType, []byte(boardString()+"1"))
				} else {
					players[i].clientConn.WriteMessage(msgType, []byte(boardString()+"0"))
				}
			}
			// Send won logic
			if winner != 0 {
				fmt.Println("Game Over")
				for i := 0; i < len(players); i++ {
					if i == winner-1 {
						players[i].clientConn.WriteMessage(msgType, []byte("WINNER YOU"))
					} else if winner == -1 {
						players[i].clientConn.WriteMessage(msgType, []byte("WINNER DRAW"))
					} else {
						players[i].clientConn.WriteMessage(msgType, []byte("WINNER OTHER"))
					}
				}
			}
		}
	}
}

func connect4Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/adamhamilton2005/Code/Go/GameServer/html/connect4.html")
}

func checkDraw() bool {
	draw := true
	for j := 0; j < boardHeight; j++ {
		for i := 0; i < boardLength; i++ {
			if board[j][i] == 0 {
				draw = false
			}
		}
	}
	if draw == false {
		return false
	}
	return true
}

func checkWinner() int {
	// Verticals
	for i := 0; i < 7; i++ {
		if (board[0][i] == 1) && (board[1][i] == 1) && (board[2][i] == 1) && (board[3][i] == 1) {
			return 1
		} else if (board[0][i] == 2) && (board[1][i] == 2) && (board[2][i] == 2) && (board[3][i] == 2) {
			return 2
		} else if (board[1][i] == 1) && (board[2][i] == 1) && (board[3][i] == 1) && (board[4][i] == 1) {
			return 1
		} else if (board[1][i] == 2) && (board[2][i] == 2) && (board[3][i] == 2) && (board[4][i] == 2) {
			return 2
		} else if (board[2][i] == 1) && (board[3][i] == 1) && (board[4][i] == 1) && (board[5][i] == 1) {
			return 1
		} else if (board[2][i] == 2) && (board[3][i] == 2) && (board[4][i] == 2) && (board[5][i] == 2) {
			return 2
		}
	}
	// Horizontals
	for j := 0; j < 6; j++ {
		if (board[j][0] == 1) && (board[j][1] == 1) && (board[j][2] == 1) && (board[j][3] == 1) {
			return 1
		} else if (board[j][0] == 2) && (board[j][1] == 2) && (board[j][2] == 2) && (board[j][3] == 2) {
			return 2
		} else if (board[j][1] == 1) && (board[j][2] == 1) && (board[j][3] == 1) && (board[j][4] == 1) {
			return 1
		} else if (board[j][1] == 2) && (board[j][2] == 2) && (board[j][3] == 2) && (board[j][4] == 2) {
			return 2
		} else if (board[j][2] == 1) && (board[j][3] == 1) && (board[j][4] == 1) && (board[j][5] == 1) {
			return 1
		} else if (board[j][2] == 2) && (board[j][3] == 2) && (board[j][4] == 2) && (board[j][5] == 2) {
			return 2
		} else if (board[j][3] == 1) && (board[j][4] == 1) && (board[j][5] == 1) && (board[j][6] == 1) {
			return 1
		} else if (board[j][3] == 2) && (board[j][4] == 2) && (board[j][5] == 2) && (board[j][6] == 2) {
			return 2
		}
	}
	// Diagonals
	if (board[3][0] == 1) && (board[2][1] == 1) && (board[1][2] == 1) && (board[0][3] == 1) {
		return 1
	} else if (board[4][0] == 1) && (board[3][1] == 1) && (board[2][2] == 1) && (board[1][3] == 1) {
		return 1
	} else if (board[3][1] == 1) && (board[2][2] == 1) && (board[1][3] == 1) && (board[0][4] == 1) {
		return 1
	} else if (board[5][0] == 1) && (board[4][1] == 1) && (board[3][2] == 1) && (board[2][3] == 1) {
		return 1
	} else if (board[4][1] == 1) && (board[3][2] == 1) && (board[2][3] == 1) && (board[1][4] == 1) {
		return 1
	} else if (board[3][2] == 1) && (board[2][3] == 1) && (board[1][4] == 1) && (board[0][5] == 1) {
		return 1
	} else if (board[5][1] == 1) && (board[4][2] == 1) && (board[3][3] == 1) && (board[2][4] == 1) {
		return 1
	} else if (board[4][2] == 1) && (board[3][3] == 1) && (board[2][4] == 1) && (board[1][5] == 1) {
		return 1
	} else if (board[3][3] == 1) && (board[2][4] == 1) && (board[1][5] == 1) && (board[0][6] == 1) {
		return 1
	} else if (board[5][2] == 1) && (board[4][3] == 1) && (board[3][4] == 1) && (board[2][5] == 1) {
		return 1
	} else if (board[4][3] == 1) && (board[3][4] == 1) && (board[2][5] == 1) && (board[1][6] == 1) {
		return 1
	} else if (board[5][3] == 1) && (board[4][4] == 1) && (board[3][5] == 1) && (board[2][6] == 1) {
		return 1
	}
	if (board[3][0] == 2) && (board[2][1] == 2) && (board[1][2] == 2) && (board[0][3] == 2) {
		return 2
	} else if (board[4][0] == 2) && (board[3][1] == 2) && (board[2][2] == 2) && (board[1][3] == 2) {
		return 2
	} else if (board[3][1] == 2) && (board[2][2] == 2) && (board[1][3] == 2) && (board[0][4] == 2) {
		return 2
	} else if (board[5][0] == 2) && (board[4][1] == 2) && (board[3][2] == 2) && (board[2][3] == 2) {
		return 2
	} else if (board[4][1] == 2) && (board[3][2] == 2) && (board[2][3] == 2) && (board[1][4] == 2) {
		return 2
	} else if (board[3][2] == 2) && (board[2][3] == 2) && (board[1][4] == 2) && (board[0][5] == 2) {
		return 2
	} else if (board[5][1] == 2) && (board[4][2] == 2) && (board[3][3] == 2) && (board[2][4] == 2) {
		return 2
	} else if (board[4][2] == 2) && (board[3][3] == 2) && (board[2][4] == 2) && (board[1][5] == 2) {
		return 2
	} else if (board[3][3] == 2) && (board[2][4] == 2) && (board[1][5] == 2) && (board[0][6] == 2) {
		return 2
	} else if (board[5][2] == 2) && (board[4][3] == 2) && (board[3][4] == 2) && (board[2][5] == 2) {
		return 2
	} else if (board[4][3] == 2) && (board[3][4] == 2) && (board[2][5] == 2) && (board[1][6] == 2) {
		return 2
	} else if (board[5][3] == 2) && (board[4][4] == 2) && (board[3][5] == 2) && (board[2][6] == 2) {
		return 2
	}
	if (board[2][0] == 1) && (board[3][1] == 1) && (board[4][2] == 1) && (board[5][3] == 1) {
		return 1
	} else if (board[1][0] == 1) && (board[2][1] == 1) && (board[3][2] == 1) && (board[4][3] == 1) {
		return 1
	} else if (board[2][1] == 1) && (board[3][2] == 1) && (board[4][3] == 1) && (board[5][4] == 1) {
		return 1
	} else if (board[0][0] == 1) && (board[1][1] == 1) && (board[2][2] == 1) && (board[3][3] == 1) {
		return 1
	} else if (board[1][1] == 1) && (board[2][2] == 1) && (board[3][3] == 1) && (board[4][4] == 1) {
		return 1
	} else if (board[2][2] == 1) && (board[3][3] == 1) && (board[4][4] == 1) && (board[5][5] == 1) {
		return 1
	} else if (board[0][1] == 1) && (board[1][2] == 1) && (board[2][3] == 1) && (board[3][4] == 1) {
		return 1
	} else if (board[1][2] == 1) && (board[2][3] == 1) && (board[3][4] == 1) && (board[4][5] == 1) {
		return 1
	} else if (board[2][3] == 1) && (board[3][4] == 1) && (board[4][5] == 1) && (board[5][6] == 1) {
		return 1
	} else if (board[0][2] == 1) && (board[1][3] == 1) && (board[2][4] == 1) && (board[3][5] == 1) {
		return 1
	} else if (board[1][3] == 1) && (board[2][4] == 1) && (board[3][5] == 1) && (board[4][6] == 1) {
		return 1
	} else if (board[0][3] == 1) && (board[1][4] == 1) && (board[2][5] == 1) && (board[3][6] == 1) {
		return 1
	}
	if (board[2][0] == 2) && (board[3][1] == 2) && (board[4][2] == 2) && (board[5][3] == 2) {
		return 2
	} else if (board[1][0] == 2) && (board[2][1] == 2) && (board[3][2] == 2) && (board[4][3] == 2) {
		return 2
	} else if (board[2][1] == 2) && (board[3][2] == 2) && (board[4][3] == 2) && (board[5][4] == 2) {
		return 2
	} else if (board[0][0] == 2) && (board[1][1] == 2) && (board[2][2] == 2) && (board[3][3] == 2) {
		return 2
	} else if (board[1][1] == 2) && (board[2][2] == 2) && (board[3][3] == 2) && (board[4][4] == 2) {
		return 2
	} else if (board[2][2] == 2) && (board[3][3] == 2) && (board[4][4] == 2) && (board[5][5] == 2) {
		return 2
	} else if (board[0][1] == 2) && (board[1][2] == 2) && (board[2][3] == 2) && (board[3][4] == 2) {
		return 2
	} else if (board[1][2] == 2) && (board[2][3] == 2) && (board[3][4] == 2) && (board[4][5] == 2) {
		return 2
	} else if (board[2][3] == 2) && (board[3][4] == 2) && (board[4][5] == 2) && (board[5][6] == 2) {
		return 2
	} else if (board[0][2] == 2) && (board[1][3] == 2) && (board[2][4] == 2) && (board[3][5] == 2) {
		return 2
	} else if (board[1][3] == 2) && (board[2][4] == 2) && (board[3][5] == 2) && (board[4][6] == 2) {
		return 2
	} else if (board[0][3] == 2) && (board[1][4] == 2) && (board[2][5] == 2) && (board[3][6] == 2) {
		return 2
	}
	return 0
}
