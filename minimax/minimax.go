package minimax

import (
	"main/board"
	"math"
)

// Minimax performs a minimax search and returns the evaluation score for the current state
func Minimax(depth int, isMaximizing bool, player board.Element, mockBoard *[board.BOARD_SIZE][board.BOARD_SIZE]board.Element) int {
	// If the search depth is zero or the game is over, return the evaluation of the current board state
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

// getOpponent returns the opponent's player type
func getOpponent(player board.Element) board.Element {
	if player == board.TRIANGLE {
		return board.CIRCLE
	}
	return board.TRIANGLE
}
