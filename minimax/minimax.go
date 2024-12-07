package minimax

import (
	"main/board"
	"math"
)

var MockBoard [7][7]board.Element

var MockTurnCounter = 0
var MockRoundCounter = 0
var MockCurrentPlayer = board.EMPTY
var MockCircleNum = 0
var MockTriangleNum = 0
var MockFirstMove = true

func MoveThePiece(fromX, fromY, X, Y int) bool {
	if MockCurrentPlayer == board.EMPTY {
		switch MockBoard[fromX][fromY] {
		case board.CIRCLE:
			MockCurrentPlayer = board.CIRCLE
		case board.TRIANGLE:
			MockCurrentPlayer = board.TRIANGLE
		}
	}

	if ValidMoveCheck(fromX, fromY, X, Y) && SequentialMoveCheck(X, Y, fromX, fromY) {
		MockBoard[X][Y] = MockBoard[fromX][fromY]
		MockBoard[fromX][fromY] = board.EMPTY
		SwitchTurnControl()
		return true
	} else {
		return false
	}
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

	return true
}

var preMoveMemory = [2]int{0, 0}

func SequentialMoveCheck(X, Y, selectedX, selectedY int) bool { // can't move already played piece
	if MockFirstMove {
		preMoveMemory[0] = X
		preMoveMemory[1] = Y
		MockFirstMove = false
		return true
	}

	if preMoveMemory[0] == selectedX && preMoveMemory[1] == selectedY {
		return false
	}

	return true
}

func SwitchTurnControl() {
	if MockCircleNum == 1 && MockCurrentPlayer == board.CIRCLE {
		MockCurrentPlayer = board.TRIANGLE
		MockTurnCounter += 2
		MockRoundCounter++
		MockFirstMove = true
		return
	}

	if MockTriangleNum == 1 && MockCurrentPlayer == board.TRIANGLE {
		MockCurrentPlayer = board.CIRCLE
		MockTurnCounter += 2
		MockRoundCounter++
		MockFirstMove = true
		return
	}

	MockTurnCounter++
	MockRoundCounter++

	if MockTurnCounter%2 == 0 {
		if MockCurrentPlayer == board.TRIANGLE {
			MockCurrentPlayer = board.CIRCLE
		} else {
			MockCurrentPlayer = board.TRIANGLE
		}
		MockFirstMove = true
	}
}

// Minimax performs a minimax search and returns the evaluation score for the current state
func Minimax(depth int, isMaximizing bool, player board.Element, mockBoard *[board.BOARD_SIZE][board.BOARD_SIZE]board.Element) int {
	// If the search depth is zero or the game is over, return the evaluation of the current board state
	if depth == 0 {
		return evaluateBoard(player)
	}

	// Get all possible actions for the current player
	actions := GetPossibleActions(player)

	if isMaximizing {
		bestEval := math.MinInt32

		// Maximize the player's score
		for _, action := range actions {
			// Simulate the move
			MoveThePiece(action.FromX, action.FromY, action.ToX, action.ToY)

			// Recursively evaluate the new state
			eval := Minimax(depth-1, false, getOpponent(player), mockBoard)

			// Maximize the score
			if eval > bestEval {
				bestEval = eval
			}
		}

		return bestEval
	} else {
		bestEval := math.MaxInt32

		// Minimize the opponent's score
		for _, action := range actions {
			// Simulate the move
			MoveThePiece(action.FromX, action.FromY, action.ToX, action.ToY)

			// Recursively evaluate the new state
			eval := Minimax(depth-1, true, getOpponent(player), mockBoard)

			// Minimize the score
			if eval < bestEval {
				bestEval = eval
			}
		}

		return bestEval
	}
}

// evaluateBoard evaluates the board for the given player and returns a score
func evaluateBoard(player board.Element) int {
	// You can create a custom evaluation function here
	// For simplicity, this example gives a basic score based on the number of pieces
	if player == board.TRIANGLE {
		return MockTriangleNum
	}
	return MockCircleNum
}

// getOpponent returns the opponent's player type
func getOpponent(player board.Element) board.Element {
	if player == board.TRIANGLE {
		return board.CIRCLE
	}
	return board.TRIANGLE
}
