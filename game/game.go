package game

import (
	"main/board"
	"math"
	"sync"
	"sync/atomic"

	"github.com/gdamore/tcell/v2"
)

var preMoveMemory = []int{-1, -1}

var FirstMove = true

var TurnCounter int32 // its for per move for each player game logic.

var CurrentPlayer = board.EMPTY

var mu sync.Mutex

type Action struct {
	FromX, FromY, ToX, ToY int
}

func GetPossibleActions(entity board.Element) []Action {

	var actionList []Action

	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			if board.Board[i][j] == entity {
				//Move to right
				if ValidMoveCheck(i, j, i, j+1) {
					actionList = append(actionList, Action{i, j, i, j + 1})
				}

				//Move to left
				if ValidMoveCheck(i, j, i, j-1) {
					actionList = append(actionList, Action{i, j, i, j - 1})
				}

				//Move to up
				if ValidMoveCheck(i, j, i-1, j) {
					actionList = append(actionList, Action{i, j, i - 1, j})
				}

				//Move to down
				if ValidMoveCheck(i, j, i+1, j) {
					actionList = append(actionList, Action{i, j, i + 1, j})
				}
			}
		}
	}
	return actionList
}

func DeathCheck(screen tcell.Screen) {
	deathValues := deathCoordinates()
	if deathValues[0] != [2]int{-1, -1} {
		board.RemovePiece(deathValues)
		board.RenderBoard(screen, CurrentPlayer)
	}
}

func deathCoordinates() [][2]int {

	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {

			//a piece in between the different pieces, horizontal
			if j < board.BOARD_SIZE-2 {
				if board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.TRIANGLE ||
					board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.CIRCLE {
					return [][2]int{{i, j + 1}}
				}
			}

			//a piece in between the different pieces, vertical
			if i < board.BOARD_SIZE-2 {
				if board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.TRIANGLE ||
					board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.CIRCLE {
					return [][2]int{{i + 1, j}}
				}
			}

			//two pieces in between the different pieces, horizontal
			if j < board.BOARD_SIZE-3 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.TRIANGLE && board.Board[i][j+3] == board.CIRCLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.CIRCLE && board.Board[i][j+3] == board.TRIANGLE {
					return [][2]int{{i, j + 1}, {i, j + 2}}
				}
			}

			//two pieces in between the different pieces, vertical
			if i < board.BOARD_SIZE-3 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.TRIANGLE && board.Board[i+3][j] == board.CIRCLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.CIRCLE && board.Board[i+3][j] == board.TRIANGLE {
					return [][2]int{{i + 1, j}, {i + 2, j}}
				}
			}

			//three pieces in between the different pieces, horizontal
			if j < board.BOARD_SIZE-4 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.TRIANGLE && board.Board[i][j+3] == board.TRIANGLE && board.Board[i][j+4] == board.CIRCLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.CIRCLE && board.Board[i][j+3] == board.CIRCLE && board.Board[i][j+4] == board.TRIANGLE {
					return [][2]int{{i, j + 1}, {i, j + 2}, {i, j + 3}}
				}
			}

			//three pieces in between the different pieces, vertical
			if i < board.BOARD_SIZE-4 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.TRIANGLE && board.Board[i+3][j] == board.TRIANGLE && board.Board[i+4][j] == board.CIRCLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.CIRCLE && board.Board[i+3][j] == board.CIRCLE && board.Board[i+4][j] == board.TRIANGLE {
					return [][2]int{{i + 1, j}, {i + 2, j}, {i + 3, j}}
				}
			}

			//four pieces in between the different pieces, horizontal
			if j < board.BOARD_SIZE-5 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.TRIANGLE && board.Board[i][j+3] == board.TRIANGLE && board.Board[i][j+4] == board.TRIANGLE && board.Board[i][j+5] == board.CIRCLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.CIRCLE && board.Board[i][j+3] == board.CIRCLE && board.Board[i][j+4] == board.CIRCLE && board.Board[i][j+5] == board.TRIANGLE {
					return [][2]int{{i, j + 1}, {i, j + 2}, {i, j + 3}, {i, j + 4}}
				}
			}

			//four pieces in between the different pieces, vertical
			if i < board.BOARD_SIZE-5 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.TRIANGLE && board.Board[i+3][j] == board.TRIANGLE && board.Board[i+4][j] == board.TRIANGLE && board.Board[i+5][j] == board.CIRCLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.CIRCLE && board.Board[i+3][j] == board.CIRCLE && board.Board[i+4][j] == board.CIRCLE && board.Board[i+5][j] == board.TRIANGLE {
					return [][2]int{{i + 1, j}, {i + 2, j}, {i + 3, j}, {i + 4, j}}
				}
			}

			//********************************************************************************************************************

			//upper border four piece
			if i == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.CIRCLE && board.Board[i+3][j] == board.CIRCLE && board.Board[i+4][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.TRIANGLE && board.Board[i+3][j] == board.TRIANGLE && board.Board[i+4][j] == board.CIRCLE {
					return [][2]int{{i, j}, {i + 1, j}, {i + 2, j}, {i + 3, j}}
				}
			}

			//lower border four piece
			if i == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i-1][j] == board.CIRCLE && board.Board[i-2][j] == board.CIRCLE && board.Board[i-3][j] == board.CIRCLE && board.Board[i-4][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i-1][j] == board.TRIANGLE && board.Board[i-2][j] == board.TRIANGLE && board.Board[i-3][j] == board.TRIANGLE && board.Board[i-4][j] == board.CIRCLE {
					return [][2]int{{i, j}, {i - 1, j}, {i - 2, j}, {i - 3, j}}
				}
			}

			// left border four piece
			if j == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.CIRCLE && board.Board[i][j+3] == board.CIRCLE && board.Board[i][j+4] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.TRIANGLE && board.Board[i][j+3] == board.TRIANGLE && board.Board[i][j+4] == board.CIRCLE {
					return [][2]int{{i, j}, {i, j + 1}, {i, j + 2}, {i, j + 3}}
				}
			}

			// right border four piece
			if j == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j-1] == board.CIRCLE && board.Board[i][j-2] == board.CIRCLE && board.Board[i][j-3] == board.CIRCLE && board.Board[i][j-4] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j-1] == board.TRIANGLE && board.Board[i][j-2] == board.TRIANGLE && board.Board[i][j-3] == board.TRIANGLE && board.Board[i][j-4] == board.CIRCLE {
					return [][2]int{{i, j}, {i, j - 1}, {i, j - 2}, {i, j - 3}}
				}
			}

			//upper border three piece
			if i == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.CIRCLE && board.Board[i+3][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.TRIANGLE && board.Board[i+3][j] == board.CIRCLE {
					return [][2]int{{i, j}, {i + 1, j}, {i + 2, j}}
				}
			}

			//lower border three piece
			if i == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i-1][j] == board.CIRCLE && board.Board[i-2][j] == board.CIRCLE && board.Board[i-3][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i-1][j] == board.TRIANGLE && board.Board[i-2][j] == board.TRIANGLE && board.Board[i-3][j] == board.CIRCLE {
					return [][2]int{{i, j}, {i - 1, j}, {i - 2, j}}
				}
			}

			// left border three piece
			if j == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.CIRCLE && board.Board[i][j+3] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.TRIANGLE && board.Board[i][j+3] == board.CIRCLE {
					return [][2]int{{i, j}, {i, j + 1}, {i, j + 2}}
				}
			}

			// right border three piece
			if j == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j-1] == board.CIRCLE && board.Board[i][j-2] == board.CIRCLE && board.Board[i][j-3] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j-1] == board.TRIANGLE && board.Board[i][j-2] == board.TRIANGLE && board.Board[i][j-3] == board.CIRCLE {
					return [][2]int{{i, j}, {i, j - 1}, {i, j - 2}}
				}
			}

			// upper border two piece
			if i == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.CIRCLE {
					return [][2]int{{i, j}, {i + 1, j}}
				}
			}

			//lower border two piece
			if i == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i-1][j] == board.CIRCLE && board.Board[i-2][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i-1][j] == board.TRIANGLE && board.Board[i-2][j] == board.CIRCLE {
					return [][2]int{{i, j}, {i - 1, j}}
				}
			}

			// left border two piece
			if j == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.CIRCLE {
					return [][2]int{{i, j}, {i, j + 1}}
				}
			}

			// right border two piece
			if j == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j-1] == board.CIRCLE && board.Board[i][j-2] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j-1] == board.TRIANGLE && board.Board[i][j-2] == board.CIRCLE {
					return [][2]int{{i, j}, {i, j - 1}}
				}
			}

			// upper border one piece
			if i == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.CIRCLE {
					return [][2]int{{i, j}}
				}
			}

			//lower border one piece
			if i == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i-1][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i-1][j] == board.CIRCLE {
					return [][2]int{{i, j}}
				}
			}

			// left border one piece
			if j == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.CIRCLE {
					return [][2]int{{i, j}}
				}
			}

			// right border one piece
			if j == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j-1] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j-1] == board.CIRCLE {
					return [][2]int{{i, j}}
				}
			}
		}
	}

	return [][2]int{{-1, -1}}
}

func MoveThePiece(fromX, fromY, X, Y int, screen tcell.Screen) bool {
	if !ValidMoveCheck(fromX, fromY, X, Y) {
		return false
	}

	if atomic.LoadInt32(&TurnCounter) == 1 {
		if preMoveMemory[0] == fromX && preMoveMemory[1] == fromY {
			return false
		}
	}

	mu.Lock()

	board.MovePiece(fromX, fromY, X, Y, screen)

	if atomic.LoadInt32(&TurnCounter) == 0 {
		preMoveMemory = []int{X, Y}
	}

	if CurrentPlayer == board.CIRCLE && board.CircleNum == 1 ||
		CurrentPlayer == board.TRIANGLE && board.TriangleNum == 1 {
		atomic.AddInt32(&TurnCounter, 1)
	}

	atomic.AddInt32(&TurnCounter, 1)

	mu.Unlock()

	if atomic.LoadInt32(&TurnCounter) >= 2 {
		atomic.StoreInt32(&TurnCounter, 0)
		mu.Lock()
		preMoveMemory = []int{-1, -1}

		if CurrentPlayer == board.CIRCLE {
			CurrentPlayer = board.TRIANGLE

		} else {
			CurrentPlayer = board.CIRCLE
		}
		mu.Unlock()
	}

	board.MoveCounter++

	board.RenderBoard(screen, CurrentPlayer)

	return true
}

func ValidMoveCheck(fromX, fromY, X, Y int) bool { // checks location-wise valid move
	var targetDist = math.Abs(float64(X-fromX)) + math.Abs(float64(Y-fromY))

	if Y < 0 || Y >= board.BOARD_SIZE || X < 0 || X >= board.BOARD_SIZE {
		return false // Out of bounds
	}

	if board.Board[fromX][fromY] == board.EMPTY || board.Board[X][Y] != board.EMPTY || targetDist >= 2 {
		return false // Invalid move
	}

	mu.Lock()
	if board.Board[fromX][fromY] == board.CIRCLE && CurrentPlayer == board.TRIANGLE ||
		board.Board[fromX][fromY] == board.TRIANGLE && CurrentPlayer == board.CIRCLE {
		mu.Unlock()
		return false
	}
	mu.Unlock()

	return true
}

func ValidSelectCheck(X, Y int) bool {
	return board.Board[X][Y] != board.EMPTY && board.Board[X][Y] == CurrentPlayer
}

func GameOverCheck(screen tcell.Screen) int {
	if board.CircleNum == 0 {
		return 0
	}

	if board.TriangleNum == 0 {
		return 1
	}

	if board.MoveCounter == 50 {
		return -1
	}

	return 2
}
