package minimax

import (
	"main/board"
)

const (
	borderPenalty        = -50
	cornerPenalty        = -150 // More significant for corner
	pieceCountMultiplier = 1000 // Balanced weight for piece count
	centerControlBonus   = 200  // Add more reward for center control
)

func EvaluateBoard(player board.Element, boardState BoardState) int {
	score := 0

	// Positional evaluation
	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			if boardState.Board[i][j] == player {
				if isOnBorder(i, j) {
					score += borderPenalty
				}
				if isOnCorner(i, j) {
					score += cornerPenalty
				}
				score += evaluateControl(i, j) // New heuristic
			}
		}
	}

	// Piece count evaluation
	if player == board.TRIANGLE {
		score += boardState.TriangleNum * pieceCountMultiplier
	} else {
		score += boardState.CircleNum * pieceCountMultiplier
	}

	// Winning condition
	if player == board.TRIANGLE && boardState.CircleNum == 0 {
		score += int(1e9) // Arbitrary large value for a win
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
