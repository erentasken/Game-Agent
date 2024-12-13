package minimax

import (
	"main/board"
	"math"
)

const (
	borderPenalty        = -120
	cornerPenalty        = -150
	centerReward         = 100
	pieceCountMultiplier = 800
	sideBySidePenalty    = -100
)

var SavedMockBoard [board.BOARD_SIZE][board.BOARD_SIZE]board.Element
var SavedMockTurnCounter int32
var SavedMockMoveCounter int
var SavedMockCurrentPlayer board.Element
var SavedMockCircleNum int
var SavedMockTriangleNum int

func EvaluateBoard(player board.Element) int {
	var score int

	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			if MockBoard[i][j] == player {
				if isOnBorder(i, j) {
					score += borderPenalty
				}
				if isOnCorner(i, j) {
					score += cornerPenalty
				}
				if isNearCenter(i, j) {
					score += centerReward
				}
			}

			if (i+1 < board.BOARD_SIZE && MockBoard[i+1][j] == player) ||
				(j+1 < board.BOARD_SIZE && MockBoard[i][j+1] == player) {
				score += sideBySidePenalty
			}
		}
	}

	if player == board.TRIANGLE {
		score += MockTriangleNum * pieceCountMultiplier
	} else {
		score += MockCircleNum * pieceCountMultiplier
	}

	return score
}

func isOnBorder(i, j int) bool {
	return i == 0 || i == board.BOARD_SIZE-1 || j == 0 || j == board.BOARD_SIZE-1
}

func isOnCorner(i, j int) bool {
	return (i == 0 && j == 0) || (i == 0 && j == board.BOARD_SIZE-1) ||
		(i == board.BOARD_SIZE-1 && j == 0) || (i == board.BOARD_SIZE-1 && j == board.BOARD_SIZE-1)
}

func isNearCenter(i, j int) bool {
	center := board.BOARD_SIZE / 2
	return abs(i-center) <= 1 && abs(j-center) <= 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Minimax(depth int, alpha int, beta int, isMaximizing bool, player board.Element) int {
	if depth == 0 {
		return EvaluateBoard(player)
	}

	actions := GetPossibleActions(player)
	if isMaximizing {
		bestEval := math.MinInt32

		for _, action := range actions {
			SaveAndMakeMove(action)

			eval := Minimax(depth-1, alpha, beta, false, getOpponent(player))

			restoreMockBoardStates()

			bestEval = Max(bestEval, eval)

			alpha = Max(alpha, bestEval)
			if beta <= alpha {
				break
			}
		}
		return bestEval
	} else {
		bestEval := math.MaxInt32

		for _, action := range actions {
			SaveAndMakeMove(action)

			eval := Minimax(depth-1, alpha, beta, true, getOpponent(player))

			restoreMockBoardStates()

			bestEval = Min(bestEval, eval)

			beta = Min(beta, bestEval)
			if beta <= alpha {
				break
			}
		}
		return bestEval
	}
}

func SaveAndMakeMove(action Action) {
	saveMockBoardStates()

	MoveThePiece(action.FromX, action.FromY, action.ToX, action.ToY)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getOpponent(player board.Element) board.Element {
	if player == board.TRIANGLE {
		return board.CIRCLE
	}
	return board.TRIANGLE
}

func saveMockBoardStates() {
	SavedMockTurnCounter = MockTurnCounter
	SavedMockMoveCounter = MockMoveCounter
	SavedMockCurrentPlayer = MockCurrentPlayer
	SavedMockCircleNum = MockCircleNum
	SavedMockTriangleNum = MockTriangleNum

	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			SavedMockBoard[i][j] = MockBoard[i][j]
		}
	}

}

func restoreMockBoardStates() {
	MockTurnCounter = SavedMockTurnCounter
	MockMoveCounter = SavedMockMoveCounter
	MockCurrentPlayer = SavedMockCurrentPlayer
	MockCircleNum = SavedMockCircleNum
	MockTriangleNum = SavedMockTriangleNum

	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			MockBoard[i][j] = SavedMockBoard[i][j]
		}
	}

}
