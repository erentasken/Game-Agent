package game

import (
	"main/board"
	"math"

	"github.com/gdamore/tcell/v2"
)

var preMoveMemory = []int{0, 0}

var firstMove = true

var turnCounter = 0 // its for per move for each player game logic.

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

func SwitchTurn(currentPlayer *board.Element) {
	if board.CircleNum == 1 && *currentPlayer == board.CIRCLE {
		*currentPlayer = board.TRIANGLE
		turnCounter += 2
		board.RoundCounter++
		firstMove = true
		return
	}

	if board.TriangleNum == 1 && *currentPlayer == board.TRIANGLE {
		*currentPlayer = board.CIRCLE
		turnCounter += 2
		board.RoundCounter++
		firstMove = true
		return
	}

	turnCounter++
	board.RoundCounter++

	if turnCounter%2 == 0 {
		if *currentPlayer == board.TRIANGLE {
			*currentPlayer = board.CIRCLE
		} else {
			*currentPlayer = board.TRIANGLE
		}
		firstMove = true
	}

}

func SequentialMoveCheck(X, Y, selectedX, selectedY int) bool { // can't move already played piece
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

func ValidMoveCheck(X, Y, targetX, targetY int) bool {
	var targetDist = math.Abs(float64(X-targetX)) + math.Abs(float64(Y-targetY))

	if targetY < 0 || targetY >= board.BOARD_SIZE || targetX < 0 || targetX >= board.BOARD_SIZE {
		return false // Out of bounds
	}

	if board.Board[targetX][targetY] == board.EMPTY || board.Board[X][Y] != board.EMPTY || targetDist >= 2 {
		return false // Invalid move
	}
	return true
}

func GameOverCheck(screen tcell.Screen) bool {
	if board.CircleNum == 0 {
		info := "TRIANGLE WINS, GAME OVER"
		for i, r := range info {
			screen.SetContent(i, board.BOARD_SIZE+2, r, nil, tcell.StyleDefault)
		}
		screen.Show()

		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					return true // Exit on ESC
				}
			}
		}
	}

	if board.TriangleNum == 0 {
		info := "CIRCLE WINS, GAME OVER"
		for i, r := range info {
			screen.SetContent(i, board.BOARD_SIZE+3, r, nil, tcell.StyleDefault)
		}
		screen.Show()

		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					return true // Exit on ESC
				}
			}
		}
	}

	if board.RoundCounter == 50 {
		info := "DRAW, GAME OVER!"
		for i, r := range info {
			screen.SetContent(i, board.BOARD_SIZE+3, r, nil, tcell.StyleDefault)
		}

		screen.Show()
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					return true // Exit on ESC
				}
			}
		}
	}

	return false
}
