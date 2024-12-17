package minimax

import (
	"main/board"
)

const (
	borderPenalty        = -50
	cornerPenalty        = -150
	pieceCountMultiplier = 1000
	centerControlBonus   = 200
)

func EvaluateBoard(player board.Element, boardState BoardState) int {
	score := 0

	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			if boardState.Board[i][j] == player {
				if isOnBorder(i, j) {
					score += borderPenalty
				}
				if isOnCorner(i, j) {
					score += cornerPenalty
				}
				score += evaluateControl(i, j)
			}
		}
	}

	if player == board.TRIANGLE {
		score += boardState.TriangleNum * pieceCountMultiplier
	} else {
		score += boardState.CircleNum * pieceCountMultiplier
	}

	if player == board.TRIANGLE && boardState.CircleNum == 0 {
		score += int(1e9)
	} else if player == board.CIRCLE && boardState.TriangleNum == 0 {
		score += int(1e9)
	}

	return score
}

func evaluateControl(i, j int) int {
	center := board.BOARD_SIZE / 2
	return -(abs(i-center) + abs(j-center)) * centerControlBonus
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isOnBorder(i, j int) bool {
	return i == 0 || i == board.BOARD_SIZE-1 || j == 0 || j == board.BOARD_SIZE-1
}

func isOnCorner(i, j int) bool {
	return (i == 0 && j == 0) || (i == 0 && j == board.BOARD_SIZE-1) ||
		(i == board.BOARD_SIZE-1 && j == 0) || (i == board.BOARD_SIZE-1 && j == board.BOARD_SIZE-1)
}
