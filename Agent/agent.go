package agent

import (
	"main/board"
	"main/game"
	"main/minimax"
	"math"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

func AgentAction(screen tcell.Screen, element board.Element) {
	// Set the depth of the minimax search (increase for harder AI)
	const depth = 3

	// Set mock data for Minimax evaluation
	minimax.MockTurnCounter = game.TurnCounter
	minimax.MockRoundCounter = board.MoveCounter
	minimax.MockCurrentPlayer = game.CurrentPlayer
	minimax.MockCircleNum = board.CircleNum
	minimax.MockTriangleNum = board.TriangleNum
	minimax.MockFirstMove = game.FirstMove

	// Deep copy the board state to simulate moves
	deepCopyBoard(board.Board)

	actions := minimax.GetPossibleActions(element)
	if len(actions) == 0 {
		return // No possible actions
	}

	var bestAction minimax.Action
	var bestActionList []minimax.Action
	bestEval := math.MinInt32

	// Evaluate all actions using Minimax
	for _, action := range actions {
		// Simulate the move
		minimax.MoveThePiece(action.FromX, action.FromY, action.ToX, action.ToY)

		// Get the evaluation score for this move
		eval := minimax.Minimax(depth-1, false, element, &minimax.MockBoard)

		// Track the best actions
		if eval > bestEval {
			bestEval = eval
			bestActionList = []minimax.Action{action} // Reset best actions list
		} else if eval == bestEval {
			bestActionList = append(bestActionList, action) // Add to best actions list
		}

	}

	// Choose the best action
	bestAction = bestActionList[rand.Intn(len(bestActionList))]

	time.Sleep(500 * time.Millisecond)
	game.MoveThePiece(bestAction.FromX, bestAction.FromY, bestAction.ToX, bestAction.ToY, screen)

	AgentAction(screen, element)
}

// deepCopyBoard creates a new deep copy of the board to simulate moves
func deepCopyBoard(originalBoard [board.BOARD_SIZE][board.BOARD_SIZE]board.Element) {
	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			minimax.MockBoard[i][j] = originalBoard[i][j]
		}
	}
}
