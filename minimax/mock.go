package minimax

import (
	"main/board"
	"math"
	"sync/atomic"
)

var MockBoard [7][7]board.Element

var MockTurnCounter int32
var MockRoundCounter = 0
var MockCurrentPlayer = board.EMPTY
var MockCircleNum = 0
var MockTriangleNum = 0
var MockFirstMove = true

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

	if MockTurnCounter == 2 {
		MockTurnCounter = 0

		preMoveMemory = []int{-1, -1}

		if MockCurrentPlayer == board.CIRCLE {
			MockCurrentPlayer = board.TRIANGLE
		} else {
			MockCurrentPlayer = board.CIRCLE
		}
	}

	return true
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
