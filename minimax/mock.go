package minimax

import (
	"main/board"
	"math"
	"sync/atomic"
)

var MockBoard [7][7]board.Element

var MockTurnCounter int32
var MockMoveCounter = 0
var MockCurrentPlayer = board.EMPTY
var MockCircleNum = 0
var MockTriangleNum = 0

var preMoveMemory = []int{-1, -1}

func MoveThePiece(fromX, fromY, X, Y int) bool {
	if !ValidMoveCheck(fromX, fromY, X, Y) {
		return false
	}

	if MockTurnCounter == 1 {
		if preMoveMemory[0] == fromX && preMoveMemory[1] == fromY {
			return false
		}
	}

	MockBoard[X][Y] = MockBoard[fromX][fromY]
	MockBoard[fromX][fromY] = board.EMPTY

	if atomic.LoadInt32(&MockTurnCounter) == 0 {
		preMoveMemory = []int{X, Y}
	}

	if MockCurrentPlayer == board.CIRCLE && MockCircleNum == 1 ||
		MockCurrentPlayer == board.TRIANGLE && MockTriangleNum == 1 {
		MockTurnCounter++
	}

	MockTurnCounter++

	if MockTurnCounter >= 2 {
		MockTurnCounter = 0

		preMoveMemory = []int{-1, -1}

		if MockCurrentPlayer == board.CIRCLE {
			MockCurrentPlayer = board.TRIANGLE
		} else {
			MockCurrentPlayer = board.CIRCLE
		}
	}

	MockMoveCounter++

	DeathCheck()

	return true
}

func DeathCheck() {
	deathValues := deathCoordinates()

	for _, v := range deathValues {
		if v == [2]int{-1, -1} {
			continue
		}

		if MockBoard[v[0]][v[1]] == board.TRIANGLE {
			MockTriangleNum--
		} else {
			MockCircleNum--
		}

		MockBoard[v[0]][v[1]] = board.EMPTY
	}
}

func deathCoordinates() [][2]int {

	var result [][2]int // Initialize the result list to store coordinates

	result = append(result, [2]int{-1, -1})

	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {

			//a piece in between the different pieces, horizontal
			if j < board.BOARD_SIZE-2 {
				if MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.TRIANGLE ||
					MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.CIRCLE {
					result = append(result, [2]int{i, j + 1})
				}
			}

			//a piece in between the different pieces, vertical
			if i < board.BOARD_SIZE-2 {
				if MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.CIRCLE {
					result = append(result, [2]int{i + 1, j})
				}
			}

			//two pieces in between the different pieces, horizontal
			if j < board.BOARD_SIZE-3 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.TRIANGLE && MockBoard[i][j+3] == board.CIRCLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.CIRCLE && MockBoard[i][j+3] == board.TRIANGLE {
					result = append(result, [2]int{i, j + 1}, [2]int{i, j + 2})
				}
			}

			//two pieces in between the different pieces, vertical
			if i < board.BOARD_SIZE-3 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.TRIANGLE && MockBoard[i+3][j] == board.CIRCLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.CIRCLE && MockBoard[i+3][j] == board.TRIANGLE {
					result = append(result, [2]int{i + 1, j}, [2]int{i + 2, j})
				}
			}

			//three pieces in between the different pieces, horizontal
			if j < board.BOARD_SIZE-4 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.TRIANGLE && MockBoard[i][j+3] == board.TRIANGLE && MockBoard[i][j+4] == board.CIRCLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.CIRCLE && MockBoard[i][j+3] == board.CIRCLE && MockBoard[i][j+4] == board.TRIANGLE {
					result = append(result, [2]int{i, j + 1}, [2]int{i, j + 2}, [2]int{i, j + 3})
				}
			}

			//three pieces in between the different pieces, vertical
			if i < board.BOARD_SIZE-4 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.TRIANGLE && MockBoard[i+3][j] == board.TRIANGLE && MockBoard[i+4][j] == board.CIRCLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.CIRCLE && MockBoard[i+3][j] == board.CIRCLE && MockBoard[i+4][j] == board.TRIANGLE {
					result = append(result, [2]int{i + 1, j}, [2]int{i + 2, j}, [2]int{i + 3, j})
				}
			}

			//four pieces in between the different pieces, horizontal
			if j < board.BOARD_SIZE-5 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.TRIANGLE && MockBoard[i][j+3] == board.TRIANGLE && MockBoard[i][j+4] == board.TRIANGLE && MockBoard[i][j+5] == board.CIRCLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.CIRCLE && MockBoard[i][j+3] == board.CIRCLE && MockBoard[i][j+4] == board.CIRCLE && MockBoard[i][j+5] == board.TRIANGLE {
					result = append(result, [2]int{i, j + 1}, [2]int{i, j + 2}, [2]int{i, j + 3}, [2]int{i, j + 4})
				}
			}

			//four pieces in between the different pieces, vertical
			if i < board.BOARD_SIZE-5 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.TRIANGLE && MockBoard[i+3][j] == board.TRIANGLE && MockBoard[i+4][j] == board.TRIANGLE && MockBoard[i+5][j] == board.CIRCLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.CIRCLE && MockBoard[i+3][j] == board.CIRCLE && MockBoard[i+4][j] == board.CIRCLE && MockBoard[i+5][j] == board.TRIANGLE {
					result = append(result, [2]int{i + 1, j}, [2]int{i + 2, j}, [2]int{i + 3, j}, [2]int{i + 4, j})
				}
			}

			//********************************************************************************************************************

			//upper border four piece
			if i == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.CIRCLE && MockBoard[i+3][j] == board.CIRCLE && MockBoard[i+4][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.TRIANGLE && MockBoard[i+3][j] == board.TRIANGLE && MockBoard[i+4][j] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i + 1, j}, [2]int{i + 2, j}, [2]int{i + 3, j})
				}
			}

			//lower border four piece
			if i == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i-1][j] == board.CIRCLE && MockBoard[i-2][j] == board.CIRCLE && MockBoard[i-3][j] == board.CIRCLE && MockBoard[i-4][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i-1][j] == board.TRIANGLE && MockBoard[i-2][j] == board.TRIANGLE && MockBoard[i-3][j] == board.TRIANGLE && MockBoard[i-4][j] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i - 1, j}, [2]int{i - 2, j}, [2]int{i - 3, j})
				}
			}

			// left border four piece
			if j == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.CIRCLE && MockBoard[i][j+3] == board.CIRCLE && MockBoard[i][j+4] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.TRIANGLE && MockBoard[i][j+3] == board.TRIANGLE && MockBoard[i][j+4] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i, j + 1}, [2]int{i, j + 2}, [2]int{i, j + 3})
				}
			}

			// right border four piece
			if j == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j-1] == board.CIRCLE && MockBoard[i][j-2] == board.CIRCLE && MockBoard[i][j-3] == board.CIRCLE && MockBoard[i][j-4] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j-1] == board.TRIANGLE && MockBoard[i][j-2] == board.TRIANGLE && MockBoard[i][j-3] == board.TRIANGLE && MockBoard[i][j-4] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i, j - 1}, [2]int{i, j - 2}, [2]int{i, j - 3})
				}
			}

			//upper border three piece
			if i == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.CIRCLE && MockBoard[i+3][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.TRIANGLE && MockBoard[i+3][j] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i + 1, j}, [2]int{i + 2, j})
				}
			}

			//lower border three piece
			if i == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i-1][j] == board.CIRCLE && MockBoard[i-2][j] == board.CIRCLE && MockBoard[i-3][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i-1][j] == board.TRIANGLE && MockBoard[i-2][j] == board.TRIANGLE && MockBoard[i-3][j] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i - 1, j}, [2]int{i - 2, j})
				}
			}

			// left border three piece
			if j == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.CIRCLE && MockBoard[i][j+3] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.TRIANGLE && MockBoard[i][j+3] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i, j + 1}, [2]int{i, j + 2})
				}
			}

			// right border three piece
			if j == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j-1] == board.CIRCLE && MockBoard[i][j-2] == board.CIRCLE && MockBoard[i][j-3] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j-1] == board.TRIANGLE && MockBoard[i][j-2] == board.TRIANGLE && MockBoard[i][j-3] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i, j - 1}, [2]int{i, j - 2})
				}
			}

			// upper border two piece
			if i == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i + 1, j})
				}
			}

			//lower border two piece
			if i == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i-1][j] == board.CIRCLE && MockBoard[i-2][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i-1][j] == board.TRIANGLE && MockBoard[i-2][j] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i - 1, j})
				}
			}

			// left border two piece
			if j == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i, j + 1})
				}
			}

			// right border two piece
			if j == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j-1] == board.CIRCLE && MockBoard[i][j-2] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j-1] == board.TRIANGLE && MockBoard[i][j-2] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i, j - 1})
				}
			}

			// upper border one piece
			if i == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.CIRCLE {
					result = append(result, [2]int{i, j})
				}
			}

			//lower border one piece
			if i == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i-1][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i-1][j] == board.CIRCLE {
					result = append(result, [2]int{i, j})
				}
			}

			// left border one piece
			if j == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.CIRCLE {
					result = append(result, [2]int{i, j})
				}
			}

			// right border one piece
			if j == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j-1] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j-1] == board.CIRCLE {
					result = append(result, [2]int{i, j})
				}
			}
		}
	}

	return result
}

type Action struct {
	FromX, FromY, ToX, ToY int
}

func GetPossibleActions(entity board.Element) []Action {
	var actionList []Action

	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			if MockBoard[i][j] == entity {
				// Move to right
				if ValidMoveCheck(i, j, i, j+1) {
					actionList = append(actionList, Action{i, j, i, j + 1})
				}

				// Move to left
				if ValidMoveCheck(i, j, i, j-1) {
					actionList = append(actionList, Action{i, j, i, j - 1})
				}

				// Move up
				if ValidMoveCheck(i, j, i-1, j) {
					actionList = append(actionList, Action{i, j, i - 1, j})
				}

				// Move down
				if ValidMoveCheck(i, j, i+1, j) {
					actionList = append(actionList, Action{i, j, i + 1, j})
				}
			}
		}
	}
	return actionList
}

func ValidMoveCheck(fromX, fromY, X, Y int) bool { // checks location-wise valid move
	var targetDist = math.Abs(float64(X-fromX)) + math.Abs(float64(Y-fromY))

	if Y < 0 || Y >= board.BOARD_SIZE || X < 0 || X >= board.BOARD_SIZE {
		return false // Out of bounds
	}

	if MockBoard[fromX][fromY] == board.EMPTY || MockBoard[X][Y] != board.EMPTY || targetDist >= 2 {
		return false // Invalid move
	}

	if MockBoard[fromX][fromY] == board.CIRCLE && MockCurrentPlayer == board.TRIANGLE ||
		MockBoard[fromX][fromY] == board.TRIANGLE && MockCurrentPlayer == board.CIRCLE {
		return false
	}

	return true
}
