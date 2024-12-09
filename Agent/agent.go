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

var status bool

func AgentAction(screen tcell.Screen, element board.Element) {
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

	actions := minimax.GetPossibleActions(element)
	if len(actions) == 0 {
		return // No possible actions
	}

	var bestAction minimax.Action
	var bestActionList []minimax.Action
	var evalList []int
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

		evalList = append(evalList, eval)
	}

	// Choose the best action
	bestAction = bestActionList[rand.Intn(len(bestActionList))]
	time.Sleep(500 * time.Millisecond)
	game.MoveThePiece(bestAction.FromX, bestAction.FromY, bestAction.ToX, bestAction.ToY, screen)

	// Choose another action if required
	var otherAction minimax.Action
	if len(bestActionList) >= 2 {
		// If there are multiple best actions, pick another from the list
		for _, action := range bestActionList {
			if action != bestAction {
				otherAction = action
				break
			}
		}
	} else {
		// Otherwise, pick a random action, ensuring it differs from the best action
		timeout := time.After(4 * time.Second)
		timedOut := false

		for {
			select {
			case <-timeout:
				// Timeout fallback: choose the first action different from the best action
				for _, action := range actions {
					if action != bestAction && otherAction.FromX != bestAction.ToX || otherAction.FromY != bestAction.ToY {
						otherAction = action
						break
					}
				}
				timedOut = true
			default:
				otherAction = actions[rand.Intn(len(actions))]
				if otherAction.FromX != bestAction.FromX || otherAction.FromY != bestAction.FromY &&
					otherAction.FromX != bestAction.ToX || otherAction.FromY != bestAction.ToY {
					break
				}
			}
			if timedOut || (otherAction.FromX != bestAction.FromX || otherAction.FromY != bestAction.FromY) {
				break
			}
		}
	}

	// Execute the additional move
	time.Sleep(500 * time.Millisecond)
	game.MoveThePiece(otherAction.FromX, otherAction.FromY, otherAction.ToX, otherAction.ToY, screen)

}

// deepCopyBoard creates a new deep copy of the board to simulate moves
func deepCopyBoard(originalBoard [board.BOARD_SIZE][board.BOARD_SIZE]board.Element) {
	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			minimax.MockBoard[i][j] = originalBoard[i][j]
		}
	}
}
