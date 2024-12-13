package minimax

import (
	"main/board"
	"math"
)

func EvaluateBoard(player board.Element) int {
	var score int

	if player == board.TRIANGLE {
		score += MockTriangleNum - MockCircleNum
	} else {
		score += MockCircleNum - MockTriangleNum
	}

	if MockMoveCounter >= 50 {
		return 0
	}

	return score
}

func Minimax(depth int, alpha int, beta int, isMaximizing bool, player board.Element) int {
	if depth == 0 {
		return EvaluateBoard(player)
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
			eval := Minimax(depth-1, alpha, beta, false, getOpponent(player))

			// Maximize the score
			if eval > bestEval {
				bestEval = eval
			}

			// Update alpha
			alpha = Max(alpha, eval)

			// Prune branches
			if beta <= alpha {
				break
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
			eval := Minimax(depth-1, alpha, beta, true, getOpponent(player))

			// Minimize the score
			if eval < bestEval {
				bestEval = eval
			}

			// Update beta
			beta = Min(beta, eval)

			// Prune branches
			if beta <= alpha {
				break
			}
		}

		return bestEval
	}
}

// Utility functions for Max and Min
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

// getOpponent returns the opponent's player type
func getOpponent(player board.Element) board.Element {
	if player == board.TRIANGLE {
		return board.CIRCLE
	}
	return board.TRIANGLE
}
