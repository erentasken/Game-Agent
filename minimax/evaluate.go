package minimax

import (
	"main/board"
)

const (
	borderPenalty        = 100
	cornerPenalty        = 150
	pieceCountMultiplier = 500
)

func EvaluateBoard(player board.Element, boardState BoardState) int {
	score := 0

	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			if boardState.Board[i][j] == getOpponent(player) {
				if isOnBorder(i, j) {
					score += borderPenalty
				}
				if isOnCorner(i, j) {
					score += cornerPenalty
				}
			}
		}
	}

	if player == board.TRIANGLE {
		score += (boardState.TriangleNum - boardState.CircleNum) * pieceCountMultiplier
	}

	if player == board.CIRCLE {
		score += (boardState.CircleNum - boardState.TriangleNum) * pieceCountMultiplier
	}

	if player == board.TRIANGLE && boardState.CircleNum == 0 {
		score += int(1e9)
	} else if player == board.CIRCLE && boardState.TriangleNum == 0 {
		score += int(1e9)
	}

	if boardState.MoveCounter == 50 {
		score = 0
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
