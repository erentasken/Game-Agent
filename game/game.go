package game

import (
	"main/board"
	"math"

	"github.com/gdamore/tcell/v2"
)

var preMoveMemory = []int{0, 0}

var firstMove = true

var turnCounter = 0 // its for per move for each player game logic.

var CurrentPlayer = board.TRIANGLE

func DeathCheck() [][2]int {

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

func MoveThePiece(fromX, fromY, X, Y int, screen tcell.Screen) {
	if ValidMoveCheck(fromX, fromY, X, Y) && sequentialMoveCheck(X, Y, fromX, fromY) {
		board.MovePiece(fromX, fromY, X, Y, screen)
		switchTurnControl()
	}
}

func ValidMoveCheck(fromX, fromY, X, Y int) bool { // checks location-wise valid move
	var targetDist = math.Abs(float64(X-fromX)) + math.Abs(float64(Y-fromY))

	if Y < 0 || Y >= board.BOARD_SIZE || X < 0 || X >= board.BOARD_SIZE {
		return false // Out of bounds
	}

	if board.Board[fromX][fromY] == board.EMPTY || board.Board[X][Y] != board.EMPTY || targetDist >= 2 {
		return false // Invalid move
	}
	return true
}

func sequentialMoveCheck(X, Y, selectedX, selectedY int) bool { // can't move already played piece
	if firstMove {
		preMoveMemory[0] = X
		preMoveMemory[1] = Y
		firstMove = false
		return true
	}

	if preMoveMemory[0] == selectedX && preMoveMemory[1] == selectedY {
		return false
	}

	return true
}

func switchTurnControl() {
	if board.CircleNum == 1 && CurrentPlayer == board.CIRCLE {
		CurrentPlayer = board.TRIANGLE
		turnCounter += 2
		board.RoundCounter++
		firstMove = true
		return
	}

	if board.TriangleNum == 1 && CurrentPlayer == board.TRIANGLE {
		CurrentPlayer = board.CIRCLE
		turnCounter += 2
		board.RoundCounter++
		firstMove = true
		return
	}

	turnCounter++
	board.RoundCounter++

	if turnCounter%2 == 0 {
		if CurrentPlayer == board.TRIANGLE {
			CurrentPlayer = board.CIRCLE
		} else {
			CurrentPlayer = board.TRIANGLE
		}
		firstMove = true
	}
}

func ValidSelectCheck(X, Y int) bool {
	return board.Board[X][Y] != board.EMPTY && board.Board[X][Y] == CurrentPlayer
}

func GameOverCheck(screen tcell.Screen) bool {
	if board.CircleNum == 0 {
		board.EndGameDisplay("Triangle", screen)
		return true
	}

	if board.TriangleNum == 0 {
		board.EndGameDisplay("Circle", screen)
		return true
	}

	if board.RoundCounter == 50 {
		board.EndGameDisplay("Draw", screen)
		return true
	}

	return false
}
