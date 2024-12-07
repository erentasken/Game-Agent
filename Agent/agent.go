package agent

import (
	"main/board"
	"main/game"
	"main/minimax"
	"math"

	"github.com/gdamore/tcell/v2"
)

var status bool

func AgentAction(screen tcell.Screen, element board.Element) {
	var preAction minimax.Action

	for i := 0; i < 2; i++ {
		// If Circle is the current player and there's only 1 circle, skip this iteration
		if game.CurrentPlayer == board.CIRCLE && board.CircleNum == 1 {
			i = 1
		}

		// Set the depth of the minimax search (increase for harder AI)
		const depth = 3

		// Set mock data for Minimax evaluation
		minimax.MockTurnCounter = game.TurnCounter
		minimax.MockRoundCounter = board.RoundCounter
		minimax.MockCurrentPlayer = game.CurrentPlayer
		minimax.MockCircleNum = board.CircleNum
		minimax.MockTriangleNum = board.TriangleNum
		minimax.MockFirstMove = game.FirstMove

		// Deep copy the board state to simulate moves
		deepCopyBoard(board.Board)

		// Get all possible actions for the agent
		actions := minimax.GetPossibleActions(element)
		if len(actions) == 0 {
			return
		}

		var bestAction minimax.Action
		bestEval := math.MinInt32

		// Loop through all actions and find the best one using Minimax
		for _, action := range actions {
			// Skip if the current action is moving the same piece as the last move
			if action.FromX == preAction.FromX && action.FromY == preAction.FromY {
				continue
			}

			// Simulate the move
			minimax.MoveThePiece(action.FromX, action.FromY, action.ToX, action.ToY)

			// Evaluate this move using Minimax
			eval := minimax.Minimax(depth-1, false, element, &minimax.MockBoard)

			// If this move is better than the previous best, store it
			if eval > bestEval {
				bestEval = eval
				bestAction = action
			}
		}

		// Perform the best action
		status := game.MoveThePiece(bestAction.FromX, bestAction.FromY, bestAction.ToX, bestAction.ToY, screen)

		// Update the last action with the current one
		preAction = bestAction

		if i == 1 && status == false {
			i = 0
		}
	}
}

// deepCopyBoard creates a new deep copy of the board to simulate moves
func deepCopyBoard(originalBoard [board.BOARD_SIZE][board.BOARD_SIZE]board.Element) {
	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			minimax.MockBoard[i][j] = originalBoard[i][j]
		}
	}
}
